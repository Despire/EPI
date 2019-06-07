package strings

import (
	"bytes"
	"math"
)

func digitCount(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 1
	}

	return int(math.Log10(float64(n))) + 1
}

func IntToString(number int) string {
	prefix := ""

	if number < 0 {
		prefix = "-"
		number = -number
	}

	if number == 0 {
		return prefix + "0"
	}

	count := digitCount(number)

	digits := make([]int, 0, count)

	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}

	b := new(bytes.Buffer)
	b.WriteString(prefix)

	for i := count - 1; i >= 0; i-- {
		b.WriteByte(byte(digits[i] + '0'))
	}

	return b.String()
}

func StringToInt(s string) int {
	prefix := ""

	switch s[0] {
	case byte('+'):
		prefix = "+"
	case byte('-'):
		prefix = "-"
	}

	result := 0
	// get rid of the prefix
	if prefix != "" {
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		result *= 10
		result += int(s[i] - '0')
	}

	if prefix == "-" {
		result = -result
	}

	return result
}
