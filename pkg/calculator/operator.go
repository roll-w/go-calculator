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
	"errors"
	"math"
)

type Precedence int

const (
	Normal Precedence = iota
	Middle
	High
)

type Type int

const (
	Infix Type = iota // + - * / ..
	Function
	Suffix // !
)

type OperatorEvaluator interface {
	Evaluate(left, right float64) (float64, error)

	// Supports returns true if the evaluator supports the given operator
	Supports(operator string) bool

	Precedence() Precedence

	Type() Type
}

type OperatorEvaluatorFactory interface {
	Create(operator string) OperatorEvaluator

	IsValid(operator string) bool
}

// NewOperatorEvaluatorFactory creates a new instance of OperatorEvaluatorFactory
//
// Supports operator evaluation for:
//
// - - * / % ^ ! sqrt log sin cos tan
func NewOperatorEvaluatorFactory() OperatorEvaluatorFactory {
	operators := map[string]OperatorEvaluator{
		"+":    additionEvaluator{},
		"-":    subtractionEvaluator{},
		"*":    multiplicationEvaluator{},
		"/":    divisionEvaluator{},
		"%":    remainderEvaluator{},
		"^":    powerEvaluator{},
		"!":    factorialEvaluator{},
		"sqrt": sqrtEvaluator{},
		"log":  logarithmEvaluator{},
		"sin":  sinEvaluator{},
		"cos":  cosEvaluator{},
		"tan":  tanEvaluator{},
	}
	return &operatorEvaluatorFactory{
		evaluators: operators,
	}
}

type operatorEvaluatorFactory struct {
	// evaluators is a map of operator to evaluator
	evaluators map[string]OperatorEvaluator
}

func (f *operatorEvaluatorFactory) IsValid(operator string) bool {
	_, ok := f.evaluators[operator]
	return ok
}

func (f *operatorEvaluatorFactory) Create(operator string) OperatorEvaluator {
	return f.evaluators[operator]
}

type (
	additionEvaluator struct {
	}
	subtractionEvaluator struct {
	}
	multiplicationEvaluator struct {
	}
	divisionEvaluator struct {
	}
	remainderEvaluator struct {
	}
	powerEvaluator struct {
	}
	factorialEvaluator struct {
	}
	sqrtEvaluator struct {
	}
	logarithmEvaluator struct {
	}
	sinEvaluator struct {
	}
	cosEvaluator struct {
	}
	tanEvaluator struct {
	}
)

func (e additionEvaluator) Evaluate(left, right float64) (float64, error) {
	return left + right, nil
}

func (e additionEvaluator) Supports(operator string) bool {
	return operator == "+"
}

func (e additionEvaluator) Precedence() Precedence {
	return Normal
}

func (e additionEvaluator) Type() Type {
	return Infix
}

func (e subtractionEvaluator) Evaluate(left, right float64) (float64, error) {
	return left - right, nil
}

func (e subtractionEvaluator) Supports(operator string) bool {
	return operator == "-"
}

func (e subtractionEvaluator) Precedence() Precedence {
	return Normal
}

func (e subtractionEvaluator) Type() Type {
	return Infix
}

func (e multiplicationEvaluator) Evaluate(left, right float64) (float64, error) {
	return left * right, nil
}

func (e multiplicationEvaluator) Supports(operator string) bool {
	return operator == "*"
}

func (e multiplicationEvaluator) Precedence() Precedence {
	return Normal
}

func (e multiplicationEvaluator) Type() Type {
	return Infix
}

func (e divisionEvaluator) Evaluate(left, right float64) (float64, error) {
	if right == 0 {
		return 0, errors.New("division by zero")
	}
	return left / right, nil
}

func (e divisionEvaluator) Supports(operator string) bool {
	return operator == "/"
}

func (e divisionEvaluator) Precedence() Precedence {
	return Middle
}

func (e divisionEvaluator) Type() Type {
	return Infix
}

func (r remainderEvaluator) Evaluate(left, right float64) (float64, error) {
	return math.Mod(left, right), nil
}

func (r remainderEvaluator) Supports(operator string) bool {
	return operator == "%"
}

func (r remainderEvaluator) Precedence() Precedence {
	return Middle
}

func (r remainderEvaluator) Type() Type {
	return Infix
}

func (e powerEvaluator) Evaluate(left, right float64) (float64, error) {
	return math.Pow(left, right), nil
}

func (e powerEvaluator) Supports(operator string) bool {
	return operator == "^"
}

func (e powerEvaluator) Precedence() Precedence {
	return High
}

func (e powerEvaluator) Type() Type {
	return Infix
}

func (e factorialEvaluator) Evaluate(left, right float64) (float64, error) {
	var result float64 = 1
	for i := 1; i <= int(left); i++ {
		result *= float64(i)
	}
	return result, nil
}

func (e factorialEvaluator) Supports(operator string) bool {
	return operator == "!"
}

func (e factorialEvaluator) Precedence() Precedence {
	return High
}

func (e factorialEvaluator) Type() Type {
	return Suffix
}

func (e sqrtEvaluator) Evaluate(left, right float64) (float64, error) {
	return math.Sqrt(left), nil
}

func (e sqrtEvaluator) Supports(operator string) bool {
	return operator == "sqrt"
}

func (e sqrtEvaluator) Precedence() Precedence {
	return High
}

func (e sqrtEvaluator) Type() Type {
	return Function
}

func (e logarithmEvaluator) Evaluate(left, right float64) (float64, error) {
	return math.Log(left), nil
}

func (e logarithmEvaluator) Supports(operator string) bool {
	return operator == "log"
}

func (e logarithmEvaluator) Precedence() Precedence {
	return High

}

func (e logarithmEvaluator) Type() Type {
	return Function
}

func (e sinEvaluator) Evaluate(left, right float64) (float64, error) {
	return math.Sin(left), nil
}

func (e sinEvaluator) Supports(operator string) bool {
	return operator == "sin"
}

func (e sinEvaluator) Precedence() Precedence {
	return High
}

func (e sinEvaluator) Type() Type {
	return Function
}

func (e cosEvaluator) Supports(operator string) bool {
	return operator == "cos"
}

func (e cosEvaluator) Precedence() Precedence {
	return High
}

func (e cosEvaluator) Type() Type {
	return Function
}

func (e cosEvaluator) Evaluate(left, right float64) (float64, error) {
	return math.Cos(left), nil
}

func (e tanEvaluator) Evaluate(left, right float64) (float64, error) {
	return math.Tan(left), nil
}

func (e tanEvaluator) Supports(operator string) bool {
	return operator == "tan"
}

func (e tanEvaluator) Precedence() Precedence {
	return High
}

func (e tanEvaluator) Type() Type {
	return Function
}
