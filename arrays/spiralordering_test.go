package arrays

import (
	"reflect"
	"testing"
)

func TestSpiralOrdering(t *testing.T) {
	tests := []struct {
		in   [][]int
		want []int
	}{
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			in: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}

	for _, test := range tests {
		got := SpiralOrdering(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Ouputs differ want:%v, got:%v", test.want, got)
		}
	}
}
