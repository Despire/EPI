package arrays

import "sort"

// Alternation rearranges the slice such that
// the elements are in the order:
// x[0] <= x[1] >= x[2] <= x[3] >= x[4] ...
// Takes O(nlogn) time and O(1) space
func Alternation(x []int) {
	sort.Ints(x)
	for i := 2; i < len(x); i += 2 {
		x[i-1], x[i] = x[i], x[i-1]
	}
}

// LinearAlternation rearranges the slice such that
// the elements are in the order:
// x[0] <= x[1] >= x[2] <= x[3] >= x[4] ...
// Takes O(n) time and O(1) space
func LinearAlternation(x []int) {
	for i := 1; i < len(x); i += 1 {
		switch {
		case i%2 == 0:
			if x[i-1] < x[i] {
				x[i], x[i-1] = x[i-1], x[i]
			}
		case i%2 == 1:
			if x[i-1] > x[i] {
				x[i], x[i-1] = x[i-1], x[i]
			}
		}
	}
}
