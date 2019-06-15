package strings

import (
	"reflect"
	"testing"
)

func TestReplaceAndRemove(t *testing.T) {
	tests := []struct {
		input []rune
		size  int
		want  []rune
	}{
		{
			input: []rune{'a', 'c', 'd', 'b', 'b', 'c', 'a'},
			size:  7,
			want:  []rune{'d', 'd', 'c', 'd', 'c', 'd', 'd'},
		},

		{
			input: []rune{'a', 'b', 'a', 'c', ' '},
			size:  4,
			want:  []rune{'d', 'd', 'd', 'd', 'c'},
		},

		{
			input: []rune{'b', 'a', 'a', 'c', ' '},
			size:  4,
			want:  []rune{'d', 'd', 'd', 'd', 'c'},
		},

		{
			input: []rune{'b', 'b', 'a', 'c', 'b', 'a', 'a'},
			size:  7,
			want:  []rune{'d', 'd', 'c', 'd', 'd', 'd', 'd'},
		},
	}

	for _, test := range tests {
		ReplaceAndRemove(test.input, test.size)

		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("Unexpected result! got: %v; want:%v", test.input, test.want)
		}
	}
}
