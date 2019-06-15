package strings

import "testing"

func TestASCIIPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "A man, a plan, a canal, Panama",
			want: true,
		},

		{
			in:   "A man",
			want: false,
		},

		{
			in:   "racecar",
			want: true,
		},
	}

	for _, test := range tests {
		if ASCIIPalindrome(test.in) != test.want {
			t.Errorf("Unwanted result! got: %v, want: %v", ASCIIPalindrome(test.in), test.want)
		}
	}
}

func TestUTF8Palindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "麥納麥",
			want: true,
		},

		{
			in:   "śćaćś",
			want: true,
		},

		{
			in:   "世界你好 好你界世",
			want: true,
		},
	}

	for _, test := range tests {
		got := UTF8Palindrome(test.in)
		if got != test.want {
			t.Errorf("Unwanted result! got: %v, want: %v", got, test.want)
		}
	}
}
