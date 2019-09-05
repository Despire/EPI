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

package heaps

// SortKIncDec sorts array that contains multiple sorted inc/dec sequences
// in one sorted sequence. Takes O(n log k) time with O(k) space complexity.
// where k is the number of sorted arrays and n is the total number of elements.
func SortKIncDec(s []int) ([]int, error) {
	var runs [][]int
	pos := 0
	for {
		run := identifyNextRun(s, pos)
		if run.len == 0 {
			break
		}
		runs = append(runs, s[run.beg:run.beg+run.len])
		pos += run.len
	}
	result, err := MergeFiles(runs...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type Run struct {
	// beg index
	beg int
	// number of elements
	len int
}

// identifyNextRun finds the longest increasing/decreasing
// sequence. The decreasing sequence will be reverse to be an
// increasing sequence.
// Takes O(N) time with O(1) space complexity.
func identifyNextRun(s []int, startPos int) Run {
	run := Run{
		beg: startPos,
		len: 0,
	}

	// empty run
	if startPos >= len(s) {
		return run
	}

	next := startPos + 1

	// run with length 1
	if next == len(s) {
		run.len = next - startPos
		return run
	}

	// increasing sequence
	if s[next-1] < s[next] {
		for next < len(s) && s[next-1] < s[next] {
			next++
		}
	} else {
		// non increasing sequence
		for next < len(s) && s[next] <= s[next-1] {
			next++
		}
		reverseSlice(s[startPos:next])
	}

	run.len = next - startPos
	return run
}

func reverseSlice(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
