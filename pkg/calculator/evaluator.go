/*
 * MIT License
 *
 * Copyright (c) 2025 RollW
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type tokenType string

const (
	/*define token types*/

	number     tokenType = "NUMBER"
	operator   tokenType = "OPERATOR"
	leftParen  tokenType = "LEFT_PAREN"
	rightParen tokenType = "RIGHT_PAREN"
	eof        tokenType = "EOF"
)

type Evaluator struct {
	OperatorEvaluatorFactory OperatorEvaluatorFactory
}

type token struct {
	tokenType tokenType
	value     string
	start     uint16
	end       uint16
}

func (t token) String() string {
	return fmt.Sprintf("%s('%s')[%d-%d]", t.tokenType, t.value, t.start, t.end)
}

func (e *Evaluator) tokenize(input string) ([]token, error) {
	iterator := 0
	var tokens []token

	operatorBuilder := strings.Builder{}
	numberBuilder := strings.Builder{}

	// TODO: fix space handling
	visitNumber := func(index int) {
		if numberBuilder.Len() == 0 {
			return
		}
		curNumber := numberBuilder.String()
		numberBuilder.Reset()
		tokens = append(tokens, token{
			tokenType: number,
			value:     curNumber,
			start:     uint16(iterator),
			end:       uint16(index),
		})
		iterator = index + len(curNumber)
	}

	visitOperator := func(index int) error {
		if operatorBuilder.Len() == 0 {
			return nil
		}
		op := operatorBuilder.String()
		operatorBuilder.Reset()
		segments, err := e.symbolSegments(op, index)
		if err != nil {
			return err
		}
		tokens = append(tokens, segments...)
		iterator += len(op)
		return nil
	}

	for index, c := range input {
		cur := char(c)

		switch {
		case cur.isNumber():
			numberBuilder.WriteRune(c)
			err := visitOperator(index)
			if err != nil {
				return nil, err
			}
		case cur.isParen():
			var t tokenType
			if cur.isLeftParen() {
				t = leftParen
			} else {
				t = rightParen
			}
			visitNumber(index)
			err := visitOperator(index)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token{
				tokenType: t,
				value:     string(cur),
				start:     uint16(iterator),
				end:       uint16(index),
			})
			iterator++
		case cur == ' ':
			if numberBuilder.Len() > 0 {
				visitNumber(index)
				break
			}
			if operatorBuilder.Len() > 0 {
				err := visitOperator(index)
				if err != nil {
					return nil, err
				}
				break
			}
			iterator++
		default:
			visitNumber(index)
			operatorBuilder.WriteRune(c)
		}
	}
	visitNumber(len(input))
	err := visitOperator(len(input))
	if err != nil {
		return nil, err
	}

	err = e.validate(tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (e *Evaluator) validate(tokens []token) error {
	if len(tokens) == 0 {
		return fmt.Errorf("no tokens found")
	}
	if tokens[0].tokenType == operator {
		return fmt.Errorf("expression cannot start with an operator")
	}
	if tokens[len(tokens)-1].tokenType == operator {
		return fmt.Errorf("expression cannot end with an operator")
	}
	for i, t := range tokens {
		// Find two connected numbers without an operator between them
		// means the expression is invalid
		if t.tokenType == number {
			if i+1 < len(tokens) && tokens[i+1].tokenType == number {
				return fmt.Errorf("too much numbers without operator between them")
			}
		}
	}
	return nil
}

func (e *Evaluator) symbolSegments(op string, index int) ([]token, error) {
	if e.OperatorEvaluatorFactory.IsValid(op) {
		return []token{
			{
				tokenType: operator,
				value:     op,
				start:     uint16(index),
				end:       uint16(index),
			},
		}, nil
	}
	tokens := make([]token, 0)

	tmpIdx := index
	operatorBuilder := strings.Builder{}
	for i, c := range op {
		if c != ' ' {
			operatorBuilder.WriteRune(c)
		}
		tmpIdx = index + i
		curOp := operatorBuilder.String()
		if e.OperatorEvaluatorFactory.IsValid(curOp) {
			tokens = append(tokens, token{
				tokenType: operator,
				value:     curOp,
				start:     uint16(index),
				end:       uint16(tmpIdx),
			})
			operatorBuilder.Reset()
		} else if c == ' ' {
			return nil, fmt.Errorf("invalid space(s) in number or operator")
		}
	}

	if operatorBuilder.Len() > 0 {
		return nil, fmt.Errorf("invalid operator: %s",
			operatorBuilder.String())
	}

	return tokens, nil
}

func (e *Evaluator) toReversePolishNotation(tokens []token) ([]token, error) {
	stack := make([]token, 0)
	var result []token
	for _, t := range tokens {
		switch t.tokenType {
		case number:
			result = append(result, t)
		case operator:
			operatorEvaluator := e.OperatorEvaluatorFactory.Create(t.value)
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if top.tokenType == operator && operatorEvaluator.Precedence() <=
					e.OperatorEvaluatorFactory.Create(top.value).Precedence() {
					result = append(result, top)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, t)
		case leftParen:
			stack = append(stack, t)
		case rightParen:
			for len(stack) != 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top.tokenType == leftParen {
					break
				}
				result = append(result, top)
			}
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.tokenType == leftParen || top.tokenType == rightParen {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		result = append(result, top)
	}
	return result, nil
}

func (e *Evaluator) EvaluateExpression(expression string) (float64, error) {
	tokens, err := e.tokenize(expression)
	fmt.Println(tokens)
	if err != nil {
		return 0, err
	}
	if len(tokens) == 0 {
		return 0, fmt.Errorf("no tokens found")
	}
	polishNotation, err := e.toReversePolishNotation(tokens)
	if err != nil {
		return 0, err
	}
	var stack []float64
	for _, t := range polishNotation {
		switch t.tokenType {
		// TODO: allow negative numbers
		case number:
			num, err := parseNumber(t.value)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		case operator:
			operatorEvaluator := e.OperatorEvaluatorFactory.Create(t.value)
			switch operatorEvaluator.Type() {
			case Function: // Function like sin, sqrt, log, etc., only one operand is required
				if len(stack) < 1 {
					return 0, fmt.Errorf("invalid expression")
				}
				right := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result, err := operatorEvaluator.Evaluate(right, 0)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			case Infix:
				if len(stack) < 2 {
					return 0, fmt.Errorf("invalid expression")
				}
				right := stack[len(stack)-1]
				left := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				result, err := operatorEvaluator.Evaluate(left, right)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			case Suffix:
				if len(stack) < 1 {
					return 0, fmt.Errorf("invalid expression")
				}
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result, err := operatorEvaluator.Evaluate(left, 0)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			}
		}
	}

	return stack[0], nil
}

func parseNumber(input string) (float64, error) {
	if strings.Contains(input, ".") {
		return strconv.ParseFloat(input, 64)
	}
	atoi, err := strconv.Atoi(input)
	if err != nil {
		return 0, nil
	}
	return float64(atoi), nil
}

type char uint8

func (c char) isNumber() bool {
	return c >= '0' && c <= '9' || c == '.'
}
func (c char) isParen() bool {
	return c == '(' || c == ')'
}

func (c char) isLeftParen() bool {
	return c == '('
}

func (c char) isRightParen() bool {
	return c == ')'
}
