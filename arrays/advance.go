package arrays

// Advance advances through the array
// and returns if it is possible to
// advance to the last position.
func Advance(x []uint) bool {
	max := 0
	for i := 0; i < len(x) && max >= i; i++ {
		d := i + int(x[i])
		if d > max {
			max = d
		}
	}
	return max >= len(x)-1
}
