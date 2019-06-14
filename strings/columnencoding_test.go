package strings

import "testing"

func TestDecodeEncoding(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "A",
			want:  1,
		},

		{
			input: "AA",
			want:  27,
		},

		{
			input: "ZZ",
			want:  702,
		},
	}

	for _, test := range tests {
		got := decodeEncoding(test.input)

		if got != test.want {
			t.Errorf("Unexpected value! got: %v; want: %v", got, test.want)
		}
	}
}
