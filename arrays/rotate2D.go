package arrays

func Rotate2DArray(matrix [][]int) {
	rows := len(matrix)

	for offset := 0; offset < rows/2; offset++ {
		for i := offset; i < rows-1-offset; i++ {
			// top left with top right
			matrix[offset][i], matrix[i][rows-1-offset] = matrix[i][rows-1-offset], matrix[offset][i]

			// top left with bottom right
			matrix[offset][i], matrix[rows-1-offset][rows-1-i] = matrix[rows-1-offset][rows-1-i], matrix[offset][i]

			// top left with bottom left
			matrix[offset][i], matrix[rows-1-i][offset] = matrix[rows-1-i][offset], matrix[offset][i]
		}
	}
}
