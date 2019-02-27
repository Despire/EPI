package arrays

// DeleteDuplicates deletes duplicate elements
// and returns the number of valid elements left
// in the slice.
// Takes O(n) time and O(1) space
func DeleteDuplicates(x []int) int {
	if len(x) == 0 {
		return 0
	}
	last := 0
	for current := 1; current < len(x); current++ {
		if x[current] != x[last] {
			last += 1
			x[last] = x[current]
		}
	}
	return last + 1
}
