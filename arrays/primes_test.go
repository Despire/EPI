package arrays

import (
	"reflect"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	tests := []struct {
		in   int
		want []uint
	}{
		{0, nil},
		{1, nil},
		{2, nil},
		{3, []uint{2}},
		{7, []uint{2, 3, 5}},
		{15, []uint{2, 3, 5, 7, 11, 13}},
		{25, []uint{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{100, []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, test := range tests {
		if got := GeneratePrimes(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("want %v; got %v", test.want, got)
		}
	}
}
