package arrays

import (
	"reflect"
	"testing"
)

func TestAlternation(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{3, 0, 8, 2, 3, 1}, []int{0, 2, 1, 3, 3, 8}},
		{[]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 1, 4, 3, 5}},
		{[]int{2, 4}, []int{2, 4}},
		{[]int{2}, []int{2}},
		{nil, nil},
	}

	for _, test := range tests {
		Alternation(test.input)
		if !reflect.DeepEqual(test.want, test.input) {
			t.Errorf("want %v; got %v", test.want, test.input)
		}

		for i := 1; i < len(test.input); i += 2 {
			if !(test.input[i] >= test.input[i-1]) {
				t.Errorf("wrong alternation at index %v, should be >= than index at %v", i, i-1)
			}
		}
		for i := 2; i < len(test.input); i += 2 {
			if !(test.input[i] <= test.input[i-1]) {
				t.Errorf("wrong alternation at index %v, should be >= than index at %v", i, i-1)
			}
		}
	}
}

func TestLinearAlternation(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{3, 0, 8, 2, 3, 1}, []int{0, 8, 2, 3, 1, 3}},
		{[]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 1, 4, 3, 5}},
		{[]int{2, 4}, []int{2, 4}},
		{[]int{2}, []int{2}},
		{nil, nil},
	}

	for _, test := range tests {
		LinearAlternation(test.input)
		if !reflect.DeepEqual(test.want, test.input) {
			t.Errorf("want %v; got %v", test.want, test.input)
		}

		for i := 1; i < len(test.input); i += 2 {
			if !(test.input[i] >= test.input[i-1]) {
				t.Errorf("wrong alternation at index %v, should be >= than index at %v", i, i-1)
			}
		}
		for i := 2; i < len(test.input); i += 2 {
			if !(test.input[i] <= test.input[i-1]) {
				t.Errorf("wrong alternation at index %v, should be >= than index at %v", i, i-1)
			}
		}
	}
}
