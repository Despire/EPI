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
	"testing"
)

func TestEvaluateRPNExpression(t *testing.T) {
	tests := []struct {
		// in
		expression string
		// out
		err    error
		result int
	}{
		{expression: "1,2,+", err: nil, result: 3},
		{expression: "1,2,+,*", err: ErrInvalidExpression, result: 0},
		{expression: "1,2,+,*", err: ErrInvalidExpression, result: 0},
		{expression: "a,2,+,*", err: ErrExpectedInteger, result: 0},
		{expression: "7,6,5,-,*", err: nil, result: -7},
		{expression: "1,1,+,-2,*", err: nil, result: -4},
		{expression: "3,4,+,2,*,1,+", err: nil, result: 15},
	}

	for _, test := range tests {
		got, err := EvaluateRPNExpression(test.expression)

		if err != test.err {
			t.Errorf("error mismatch. got: %v; expected: %v", err, test.err)
		}

		if got != test.result {
			t.Errorf("result mismatch. got: %v; want: %v", got, test.result)
		}
	}
}
