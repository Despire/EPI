package arrays

import (
	"reflect"
	"testing"
)

func TestRotate2DArray(t *testing.T) {
	tests := []struct {
		in   [][]int
		want [][]int
	}{
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			in: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
		{
			in: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}
	for _, test := range tests {
		Rotate2DArray(test.in)
		if !reflect.DeepEqual(test.in, test.want) {
			t.Errorf("Ouputs differ want:%v, got:%v", test.want, test.in)
		}
	}

}
