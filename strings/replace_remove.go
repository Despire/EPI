package strings

// Time complexity: O(N)
// Space complexity: O(1)

// ReplaceAndRemove replaces each 'a' by two 'd' and removes
// each occurence of 'b'
func ReplaceAndRemove(chars []rune, size int) int {
	aCount := 0
	read := 0
	write := 0

	for read < size {
		if chars[read] == 'a' {
			aCount++
		}
		if chars[read] != 'b' {
			chars[write] = chars[read]
			write++
		}
		read++
	}

	read = write - 1
	write = write + aCount - 1
	size = write + 1
	for aCount > 0 {
		if chars[read] == 'a' {
			aCount--

			chars[write] = 'd'
			chars[write-1] = 'd'

			write--
		} else {
			chars[write] = chars[read]
		}
		read--
		write--
	}

	return size
}
