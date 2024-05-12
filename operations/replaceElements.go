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
