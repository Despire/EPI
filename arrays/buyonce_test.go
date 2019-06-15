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

import "testing"

func TestBuyAndSellStockOnce(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}, 30},
		{[]int{210, 315, 275, 295, 260, 270, 290, 230, 255, 250}, 105},
		{[]int{2, 4, 2, 6, 3, 0, 5}, 5},
		{[]int{200, 100}, 0},
		{nil, 0},
	}
	for _, test := range tests {
		got := BuyAndSellStockOnce(test.in)
		if got != test.want {
			t.Errorf("want %v; got %v", test.want, got)
		}
	}
}
