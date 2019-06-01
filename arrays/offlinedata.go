package arrays

import "math/rand"

// if we have a set of size 1 it's trivial just return the element
// if we have a set of 2 and subsetSize is 1 just return the random of the two.
// if we have a set of 3 and subsetSize is 1 just return the random of the three.
// if we have a set of 3 and subsetSize is 2 just return the random of three and then random of two.
// if we have a set of 3 and subsetSize is 3 just return the random of the first then random of the two and than the last element left.

func sampleOfflineData(data []int, subsetSize int) []int {
	if subsetSize < 1 {
		return data[:0]
	}

	for i := range data {
		if i >= subsetSize {
			break
		}
		startElement := i + rand.Intn(len(data)-i)
		data[i], data[startElement] = data[startElement], data[i]
	}

	return data[:subsetSize]
}