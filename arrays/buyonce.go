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

import "math"

// BuyAndSellStockOnce finds the maximum profit
// possible from the given stock prices in the slice
// returns the maximum profit.
// The time complexity is O(N) and the space complexity is O(1).
func BuyAndSellStockOnce(prices []int) int {
	var min, profit float64 = math.MaxFloat64, 0
	for i := 0; i < len(prices); i++ {
		min = math.Min(min, float64(prices[i]))
		profit = math.Max(profit, float64(prices[i])-min)
	}
	return int(profit)
}
