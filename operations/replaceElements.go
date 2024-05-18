/*!
* \file replaceElements.go
* \brief Functions for replacing elements in a matrix.
 */
package operations

/*!
* \brief Replaces one column of the matrix with another.
*
* \param matrix Matrix.
* \param originCol Index of the column to be replaced (1-based).
* \param destCol Index of the destination column (1-based).
 */
func replaceCol(matrix [][]int, originCol int, destCol int) {
	originCol--
	destCol--
	for i := range matrix {
		temp := matrix[i][destCol]
		matrix[i][destCol] = matrix[i][originCol]
		matrix[i][originCol] = temp
	}
}

/*!
* \brief Replaces one row of the matrix with another.
*
* \param matrix Matrix.
* \param originRow Index of the row to be replaced (1-based).
* \param destRow Index of the destination row (1-based).
 */
func replaceRow(matrix [][]int, originRow int, destRow int) {
	originRow--
	destRow--
	for i := range matrix {
		temp := matrix[destRow][i]
		matrix[destRow][i] = matrix[originRow][i]
		matrix[originRow][i] = temp
	}
}

/*!
* \brief Transposes the matrix.
*
* \param matrix Matrix to be transposed.
 */
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
