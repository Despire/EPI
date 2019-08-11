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
// limitations under the Licens

package strings

import "strconv"

// ComputeIpv4 computes all valid ip addresses from a strings
// which represents an ip address but does not contains dots (for the individual octects).
// e.g 19216801.
// Takes O(1) time and O(1) space complexity.
func ComputeIpv4(s string) []string {
	result := make([]string, 0, 1)
	l := len(s)

	// i,j,k represents the size of each part.
	for i := 1; i < 4 && i < l; i++ {

		if !isOctetValid(s[:i]) {
			continue
		}

		for j := 1; j < 4 && i+j < l; j++ {

			if !isOctetValid(s[i : i+j]) {
				continue
			}

			for k := 1; k < 4 && i+k+j < l; k++ {

				if !isOctetValid(s[i+j : i+j+k]) {
					continue
				}

				if !isOctetValid(s[i+j+k:]) {
					continue
				}

				if len(s[i+j+k:]) > 3 {
					continue
				}
				result = append(result, string(
					s[:i]+"."+s[i:i+j]+"."+s[i+j:i+j+k]+"."+s[i+j+k:],
				))
			}
		}
	}

	return result
}

// isOctetValid checks if the octect from the IPv4
// address is a valid octect.
func isOctetValid(oct string) bool {
	var value int64
	var err error

	if value, err = strconv.ParseInt(oct, 10, 64); err != nil {
		return false
	}

	return value < 255 && value > 0
}
