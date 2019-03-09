package arrays

import (
	"reflect"
	"testing"
)

func TestApplePermutation(t *testing.T) {
	tests := []struct {
		in   []int
		perm []int
		want []int
	}{
		{[]int{0, 1, 2, 3}, []int{2, 0, 1, 3}, []int{1, 2, 0, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 2, 1, 3}, []int{0, 2, 1, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, []int{3, 2, 1, 0}, []int{3, 2, 1, 0}},
		{[]int{0, 1, 2, 3}, []int{3, 1, 2, 0}, []int{3, 1, 2, 0}},
	}

	for _, test := range tests {
		ApplyPermutation(test.in, test.perm)
		if !reflect.DeepEqual(test.want, test.in) {
			t.Errorf("want %v, got %v", test.want, test.in)
		}
	}
}
