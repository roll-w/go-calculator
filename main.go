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

package main

import (
	"bufio"
	"fmt"
	"go-calculator/pkg/calculator"
	"os"
	"strconv"
	"strings"
)

func readExpression() (string, error) {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " "), nil
	}

	fmt.Print("Enter an expression: ")
	input := bufio.NewReader(os.Stdin)
	inputString, err := input.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	inputString = strings.TrimSuffix(inputString, "\n")
	inputString = strings.TrimSuffix(inputString, "\r")

	return inputString, nil
}

func main() {
	inputString, err := readExpression()
	if err != nil {
		fmt.Printf("Error reading expression: %s\n", err)
		os.Exit(1)
		return
	}

	evaluator := calculator.Evaluator{OperatorEvaluatorFactory: calculator.NewOperatorEvaluatorFactory()}
	res, err := evaluator.EvaluateExpression(inputString)
	if err != nil {
		fmt.Printf("Error evaluating expression: %s\n", err)
		os.Exit(1)
		return
	}
	formatted := strconv.FormatFloat(res, 'f', -1, 64)
	fmt.Printf("%s\n", formatted)
}
