package arrays

import (
	"reflect"
	"testing"
)

func TestIncrement(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		{[]int{9}, []int{1, 0}},
		{[]int{1, 9}, []int{2, 0}},
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		{[]int{1, 8, 9}, []int{1, 9, 0}},
		{[]int{2, 0, 1}, []int{2, 0, 2}},
	}

	for _, test := range tests {
		got := Increment(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want %v; got %v", test.want, got)
		}
	}
}
