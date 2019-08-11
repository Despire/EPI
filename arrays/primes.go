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

package arrays

// GeneratePrimes generates all prime numbers from 2
// all the way up but no including upperBound.
// Time complexity is O(n log(log(n))), space complexity O(n)
func GeneratePrimes(upperBound int) []uint {
	if upperBound < 3 {
		return nil
	}

	isPrime := make([]bool, upperBound)
	for i := range isPrime {
		isPrime[i] = true
	}
	var primes []uint
	for i := 2; i < len(isPrime); i++ {
		if isPrime[i] {
			primes = append(primes, uint(i))
			for j := i; j < len(isPrime); j += i {
				isPrime[j] = false
			}
		}
	}
	return primes
}
