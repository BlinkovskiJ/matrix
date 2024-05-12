package operations

func transpose(matrix [][]int) {
	transposedMatrix := make([][]int, len(matrix))
	for i := range matrix {
		for j := range matrix[i] {
			transposedMatrix[i] = append(transposedMatrix[i], matrix[j][i])
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = transposedMatrix[i][j]
		}
	}
}
