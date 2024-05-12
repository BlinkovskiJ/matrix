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
