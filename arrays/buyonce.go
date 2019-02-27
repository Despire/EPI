package arrays

import "math"

// BuyAndSellStockOnce finds the maximum profit
// possible from the given stock prices in the slice
// returns the maximum profit.
// Takes O(n) time and O(1) space
func BuyAndSellStockOnce(prices []int) int {
	var min, profit float64 = math.MaxFloat64, 0
	for i := 0; i < len(prices); i++ {
		min = math.Min(min, float64(prices[i]))
		profit = math.Max(profit, float64(prices[i]) - min)
	}
	return int(profit)
}
