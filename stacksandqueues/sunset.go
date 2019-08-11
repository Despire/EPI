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

// ComputeSunsetView processes building heights in
// east-to-west order and returns the set of buildings
// which view the sunset.
func ComputeSunsetView(heights []int) []int {
	var view []int

	//    <- 	WEST																		EAST ->
	//                           __________
	//   __________             |         |
	//   |        |             |         |
	//   |        |   ________  |         |
	//   |        |  |       |  |   6     |
	//   |  5     |  |       |  |         |
	//   |        |  |  3    |  |         |
	//   |        |  |       |  |         |
	//  _______________________________________________________________________________________________

	if len(heights) == 0 {
		return make([]int, 0)
	}

	var (
		stack = new(stack)
	)
	for i := 0; i < len(heights); i++ {
		currHeight := heights[i]
		// for all the previous building heights
		// that are lower than the curren building
		// remove them from the stack.
		for stack.Size() > 0 && currHeight >= heights[stack.Top().(int)] {
			stack.Pop()
		}
		stack.Push(i)
	}

	view = make([]int, 0, stack.Size())

	for stack.Size() > 0 {
		view = append(view, stack.Pop().(int))
	}

	return view
}
