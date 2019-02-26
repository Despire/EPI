package arrays

// Partitions the slice in linear time, such that all elements
// less than the pivot appears before all elements
// equal to the pivot and elements which are greater than the pivot
// appears as last.
// The pivot is selected as elems[i].
func partition(elems []int, i int) (int, int) {
	// paritions the slice in O(n) time.
	// such that elems[0:s] contains all elements that are '<p'
	// elems[s:g] contains all elements that are '=p'
	// elems[g:] contains all the elements that are '>p"
	p := elems[i]
	s := 0
	g := len(elems)
	e := 0

	for e < g {
		switch {
		case elems[e] < p:
			elems[s], elems[e] = elems[e], elems[s]
			s += 1
			e += 1
		case elems[e] > p:
			g -= 1
			elems[g], elems[e] = elems[e], elems[g]
		default:
			// if isn't '>p' '<p' than it is equal
			e += 1
		}
	}
	return s, g
}
