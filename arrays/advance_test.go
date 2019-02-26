package arrays

import "testing"

func TestAdvance(t *testing.T) {
	tests := []struct {
		in   []uint
		want bool
	}{
		{[]uint{3, 3, 1, 0, 2, 0, 1}, true},
		{[]uint{3, 2, 0, 0, 2, 0, 1}, false},
		{[]uint{3, 2, 0, 1, 2, 0, 1}, true},
		{[]uint{5, 2, 0, 0, 0, 0, 1}, false},
		{[]uint{1, 5, 0, 0, 0, 0, 1}, true},
		{[]uint{6, 0, 0, 0, 0, 0, 1}, true},
		{[]uint{1, 1, 0, 1, 1, 1, 1}, false},
		{[]uint{1, 1, 2, 0, 1, 0, 1}, false},
	}

	for _, test := range tests {
		got := Advance(test.in)
		if got != test.want {
			t.Errorf("want %v; got %v", test.want, got)
		}
	}
}
