// Copyright (c) 2019 Matúš Mrekaj.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License

package stacksandqueues

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrInvalidExpression is return if the expression passed
	// to the EvaluateRPNExpression is not valid.
	ErrInvalidExpression = errors.New("expression is not a valid RPN expression")

	// ErrExpectedInteger is returned if the expression contains
	// a invalid interger value.
	ErrExpectedInteger = errors.New("exprected a interger value")
)

// operatorCallbacks provides a callback function for
// each operator.
var operatorCallbacks = map[string]func(a, b int) int{
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
}

// EvaluateRPNExpression takes a string
// which represents an expresion in RPN (reverse polish notation)
// and returns the result of the expression.
func EvaluateRPNExpression(exp string) (int, error) {
	tokens := strings.Split(exp, ",")
	stack := new(stack)

	// either two values,
	// two operators,
	// one value, one operator,
	// which is an invalid expression.
	if len(tokens) < 2 {
		return 0, ErrInvalidExpression
	}

	for i := 0; i < 2; i++ {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			return 0, ErrExpectedInteger
		}
		stack.Push(val)
	}
	tokens = tokens[2:]

	for i, token := range tokens {
		if callback, ok := operatorCallbacks[token]; ok {
			if stack.Size() < 2 {
				return 0, ErrInvalidExpression
			}
			left := stack.Pop().(int)
			right := stack.Pop().(int)
			stack.Push(callback(left, right))
		} else {
			val, err := strconv.Atoi(tokens[i])
			if err != nil {
				return 0, ErrExpectedInteger
			}
			stack.Push(val)
		}
	}
	if stack.Size() != 1 {
		return 0, ErrInvalidExpression
	}
	result := stack.Pop().(int)

	return result, nil
}
