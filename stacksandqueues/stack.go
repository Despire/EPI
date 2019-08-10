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
	"fmt"
	"reflect"
)

type Stack interface {
	// Push pushes <value> on the top of the stack.
	Push(interface{})
	// Pop removes the first value from the top of the stack.
	Pop() interface{}
	// Size returns the number of elements in the stack.
	Size() int
	// Top returns the first element from the stack.
	Top() interface{}
}

type stack struct {
	typ      reflect.Type
	elements []interface{}
}

func (s *stack) checkType(value interface{}) {
	typ := reflect.TypeOf(value)

	if s.typ != nil && s.typ.PkgPath()+s.typ.Name() != typ.PkgPath()+typ.Name() {
		panic(
			fmt.Sprintf("pushing value with type %s, however the stack contains values with type %s", typ.PkgPath()+typ.Name(), s.typ.PkgPath()+s.typ.Name()),
		)
	}
}

func (s *stack) Push(value interface{}) {
	// will trigger only for the first value pushed
	if s.typ == nil {
		s.typ = reflect.TypeOf(value)
	}

	s.checkType(value)

	s.elements = append(s.elements, value)
}

func (s *stack) Pop() interface{} {
	if len(s.elements) == 0 {
		panic("no elements to pop")
	}
	result := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return result
}

func (s *stack) Size() int {
	return len(s.elements)
}

func (s *stack) Top() interface{} {
	if len(s.elements) == 0 {
		panic("no element to return")
	}
	return s.elements[len(s.elements)-1]
}

type IntMaxValue struct {
	count int
	value int
}

type IntStack struct {
	elements  []int
	maxValues []IntMaxValue
}

func (s *IntStack) Max() interface{} {
	if len(s.maxValues) == 0 {
		return nil
	}
	return s.maxValues[len(s.maxValues)-1].value
}

func (s *IntStack) Push(value interface{}) {
	if _, ok := value.(int); !ok {
		panic("pushed a non-integer value")
	}
	s.elements = append(s.elements, value.(int))

	if len(s.maxValues) == 0 {
		s.maxValues = append(s.maxValues, IntMaxValue{
			count: 1,
			value: value.(int),
		})
	} else {
		if s.maxValues[len(s.maxValues)-1].value == value.(int) {
			s.maxValues[len(s.maxValues)-1].count++
		} else if s.maxValues[len(s.maxValues)-1].value < value.(int) {
			s.maxValues = append(s.maxValues, IntMaxValue{
				count: 1,
				value: value.(int),
			})
		}
	}
}

func (s *IntStack) Pop() interface{} {
	if len(s.elements) == 0 {
		panic("no elements to pop")
	}
	result := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	if s.maxValues[len(s.maxValues)-1].value == result {
		s.maxValues[len(s.maxValues)-1].count--
		if s.maxValues[len(s.maxValues)-1].count == 0 {
			s.maxValues = s.maxValues[:len(s.maxValues)-1]
		}
	}
	return result
}

func (s *IntStack) Size() int {
	return len(s.elements)
}

func (s *IntStack) Top() interface{} {
	if len(s.elements) == 0 {
		panic("no elements to return")
	}
	return s.elements[len(s.elements)-1]
}
