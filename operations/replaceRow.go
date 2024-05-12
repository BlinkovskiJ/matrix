package operations

func replaceRow(matrix [][]int, originRow int, destRow int) {
	originRow--
	destRow--
	for i := range matrix {
		temp := matrix[destRow][i]
		matrix[destRow][i] = matrix[originRow][i]
		matrix[originRow][i] = temp
	}
}
