package arrays

import "math"

func SpiralOrdering(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([]int, 0)
	for offset := 0; offset < int(math.Ceil(float64(rows)/2.0)); offset++ {
		if offset == cols-1-offset {
			result = append(result, matrix[offset][offset])
			continue
		}
		for i := offset; i < cols-1-offset; i++ {
			result = append(result, matrix[offset][i])
		}
		for i := offset; i < rows-1-offset; i++ {
			result = append(result, matrix[i][cols-1-offset])
		}
		for i := cols - 1 - offset; i > offset; i-- {
			result = append(result, matrix[rows-1-offset][i])
		}
		for i := rows - 1 - offset; i > offset; i-- {
			result = append(result, matrix[i][offset])
		}
	}
	return result
}
