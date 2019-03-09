package arrays

func ApplyPermutation(elems, permutation []int) {
	s := make([]bool, len(elems))

	for i := range elems {
		current := i
		for s[i] == false {
			elems[i], elems[permutation[current]] = elems[permutation[current]], elems[i]
			s[permutation[current]] = true
			current = permutation[current]
		}
	}

}
