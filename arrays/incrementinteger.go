package arrays

// increment increments the number passed in as a slice of digits
// by one, and returns the new number.
// Time complexity O(n).
func increment(digits []int) []int {
	last := len(digits) - 1
	digits[last] += 1
	for i := last; i > 0 && digits[i] == 10; i-- {
		digits[i] = 0
		digits[i-1] += 1
	}
	if digits[0] == 10 {
		digits = append([]int{1, 0}, digits[1:]...)
	}
	return digits
}
