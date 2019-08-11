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
// limitations under the License.

package strings

// keyboardMap maps keys (keyboard) to
// characters for that specific key.
var keyboardMap = map[byte]string{
	1: "",
	2: "ABC",
	3: "DEF",
	4: "GHI",
	5: "JKL",
	6: "MNO",
	7: "PQRS",
	8: "TUV",
	9: "WXYZ",
	0: "",
}

// ComputeWordsFromNumber takes as input a phone number encoded in a string.
// and returns all possible character sequence that correspond to the phone number.
// Takes O((4^n)*n) time complexity.
func ComputeWordsFromNumber(digits string) ([]string, error) {
	word := make([]byte, len(digits))
	result := make([]string, 0, 10)
	return compute(digits, 0, word, result), nil
}

func compute(digits string, idx int, word []byte, result []string) []string {
	if idx >= len(word) {
		return append(result, string(word))
	}

	set := keyboardMap[digits[idx]-'0']

	for _, character := range set {
		word[idx] = byte(character)
		result = compute(digits, idx+1, word, result)
	}

	return result
}
