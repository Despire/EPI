package strings

import "testing"

func TestDigitCount(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 5,
			want:  1,
		},

		{
			input: 10,
			want:  2,
		},

		{
			input: 153,
			want:  3,
		},

		{
			input: 12932198,
			want:  8,
		},

		{
			input: 1209,
			want:  4,
		},
	}

	for _, test := range tests {
		got := digitCount(test.input)

		if got != test.want {
			t.Errorf("Unwanted result ! got result: %v; wanted result: %v", got, test.want)
		}
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 123,
			want:  "123",
		},

		{
			input: -123,
			want:  "-123",
		},

		{
			input: 0,
			want:  "0",
		},

		{
			input: -0,
			want:  "0",
		},

		{
			input: -231,
			want:  "-231",
		},
	}

	for _, test := range tests {
		got := IntToString(test.input)
		if got != test.want {
			t.Errorf("Unwanted result! got result: %v; wanted result: %v", got, test.want)
		}
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "+123",
			want:  123,
		},

		{
			input: "-123",
			want:  -123,
		},

		{
			input: "123",
			want:  123,
		},

		{
			input: "-12837",
			want:  -12837,
		},
	}

	for _, test := range tests {
		got := StringToInt(test.input)

		if got != test.want {
			t.Errorf("Unwanted result! got result: %v; wanted result: %v", got, test.want)
		}
	}
}
