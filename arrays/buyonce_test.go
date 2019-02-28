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
