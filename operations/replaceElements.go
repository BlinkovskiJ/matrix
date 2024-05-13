package operations

func replaceCol(matrix [][]int, originCol int, destCol int) {
	originCol--
	destCol--
	for i := range matrix {
		temp := matrix[i][destCol]
		matrix[i][destCol] = matrix[i][originCol]
		matrix[i][originCol] = temp
	}
}

func replaceRow(matrix [][]int, originRow int, destRow int) {
	originRow--
	destRow--
	for i := range matrix {
		temp := matrix[destRow][i]
		matrix[destRow][i] = matrix[originRow][i]
		matrix[originRow][i] = temp
	}
}

func transpose(matrix [][]int) {
	transposedMatrix := make([][]int, len(matrix))
	for i := range matrix {
		transposedMatrix[i] = make([]int, len(matrix[i]))
		for j := range matrix[i] {
			transposedMatrix[i][j] = matrix[j][i]
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = transposedMatrix[i][j]
		}
	}
}
