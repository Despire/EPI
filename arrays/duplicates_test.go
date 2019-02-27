package arrays

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{2, 2, 3, 4, 4, 5}, 4},
		{[]int{2, 3, 4, 4, 7, 11, 11, 11, 13}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{2, 3, 5, 5, 7, 11, 11, 11, 13}, 6},
	}
	for _, test := range tests {
		got := DeleteDuplicates(test.in)
		if got != test.want {
			t.Errorf("want %v; got %v", test.want, got)
		}
	}
}
