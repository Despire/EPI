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
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		stack  *stack
		values []interface{}
		size   int
		top    interface{}
		typ    reflect.Type
	}{
		{stack: new(stack), size: 4, values: []interface{}{3, 4, 5, 6}, top: 6, typ: reflect.TypeOf(3)},
		{stack: new(stack), size: 7, values: []interface{}{"1", "2", "3", "4", "5", "6", "7"}, top: "7", typ: reflect.TypeOf("")},
	}

	for _, test := range tests {

		for _, val := range test.values {
			test.stack.Push(val)
		}

		if test.stack.Size() != test.size {
			t.Errorf("unwanted result! got:%v; want: %v", test.stack.Size(), test.size)
		}

		top := test.stack.Top()
		typ := reflect.TypeOf(top)

		if typ.PkgPath()+typ.Name() != test.typ.PkgPath()+test.typ.Name() {
			t.Errorf("unwanted result! got:%v; want:%v", typ.PkgPath()+typ.Name(), test.typ.PkgPath()+test.typ.Name())
		}

		if test.typ != typ {
			t.Errorf("unwanted result! got:%v; want:%v", typ, test.typ)
		}

		switch top.(type) {
		case int:
			if top.(int) != test.top.(int) {
				t.Errorf("unwanted result! got:%v; wat:%v", top.(int), test.top)
			}
		case string:
			if top.(string) != test.top.(string) {
				t.Errorf("unwanted result! got:%v; wat:%v", top.(int), test.top)
			}
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		stack         *stack
		values        []interface{}
		beforePopSize int
		afterPopSize  int
		beforePopTop  interface{}
		afterPopTop   interface{}
		popValue      interface{}
		typ           reflect.Type
	}{
		{stack: new(stack), values: []interface{}{3, 4, 5, 6}, typ: reflect.TypeOf(3), beforePopSize: 4, afterPopSize: 3, beforePopTop: 6, afterPopTop: 5, popValue: 6},
		{stack: new(stack), values: []interface{}{"1", "2", "3", "4", "5", "6", "7"}, typ: reflect.TypeOf(""), beforePopSize: 7, afterPopSize: 6, beforePopTop: "7", afterPopTop: "6", popValue: "7"},
		{stack: new(stack), values: []interface{}{"1", "2", "3", "5", "6"}, typ: reflect.TypeOf(""), beforePopSize: 5, afterPopSize: 4, beforePopTop: "6", afterPopTop: "5", popValue: "6"},
	}

	for _, test := range tests {

		for _, val := range test.values {
			test.stack.Push(val)
		}

		if test.stack.Size() != test.beforePopSize {
			t.Errorf("unwanted result stack size! got:%v; want: %v", test.stack.Size(), test.beforePopSize)
		}

		top := test.stack.Top()
		typ := reflect.TypeOf(top)

		if typ.PkgPath()+typ.Name() != test.typ.PkgPath()+test.typ.Name() {
			t.Errorf("unwanted result type value! got:%v; want:%v", typ.PkgPath()+typ.Name(), test.typ.PkgPath()+test.typ.Name())
		}

		if test.typ != typ {
			t.Errorf("unwanted result type interface! got:%v; want:%v", typ, test.typ)
		}

		switch top.(type) {
		case int:
			if top.(int) != test.beforePopTop.(int) {
				t.Errorf("unwanted result int value! got:%v; wat:%v", top.(int), test.beforePopTop)
			}
		case string:
			if top.(string) != test.beforePopTop.(string) {
				t.Errorf("unwanted result string value! got:%v; wat:%v", top.(int), test.beforePopTop)
			}
		}

		popedValue := test.stack.Pop()

		if test.stack.Size() != test.afterPopSize {
			t.Errorf("unwanted result! got:%v; want: %v", test.stack.Size(), test.afterPopSize)
		}

		top = test.stack.Top()
		typ = reflect.TypeOf(top)

		if typ.PkgPath()+typ.Name() != test.typ.PkgPath()+test.typ.Name() {
			t.Errorf("unwanted result! got:%v; want:%v", typ.PkgPath()+typ.Name(), test.typ.PkgPath()+test.typ.Name())
		}

		if test.typ != typ {
			t.Errorf("unwanted result! got:%v; want:%v", typ, test.typ)
		}

		switch top.(type) {
		case int:
			if top.(int) != test.afterPopTop.(int) {
				t.Errorf("unwanted result! got:%v; want:%v", top.(int), test.afterPopTop)
			}
			if popedValue.(int) != test.popValue.(int) {
				t.Errorf("unwanted result! got:%v; want:%v", popedValue.(int), test.popValue.(int))
			}
		case string:
			if top.(string) != test.afterPopTop.(string) {
				t.Errorf("unwanted result! got:%v; want:%v", top.(int), test.afterPopTop)
			}
			if popedValue.(string) != test.popValue.(string) {
				t.Errorf("unwanted result! got:%v; want:%v", popedValue.(string), test.popValue.(string))
			}
		}
	}
}

func TestMaxAPI(t *testing.T) {
	tests := []struct {
		in   interface{}
		want interface{}
	}{
		{nil, nil},
		{4, 4},
		{3, 4},
		{1, 4},
		{4, 4},
		{5, 5},
		{5, 5},
		{0, 5},
		{9, 9},
		{6, 9},
		{8, 9},
		{7, 9},
	}
	s := new(IntStack)
	for _, test := range tests {
		if test.in != nil {
			s.Push(test.in)
		}
		if got := s.Max(); got != test.want {
			t.Errorf("%+v got s.Max() = %v; want %v", s, got, test.want)
		}
	}

	for s.Size() > 0 {
		if got, want := s.Max(), tests[s.Size()].want; got != want {
			t.Errorf("%+v got s.Max() = %v; want %v", s, got, want)
		}
		s.Pop()
	}
}
