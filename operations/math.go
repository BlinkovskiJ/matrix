/*!
* \file math.go
* \brief Functions for matrix operations.
 */
package operations

/*!
* \brief Multiplies two matrices.
*
* \param matrix1 First matrix.
* \param matrix2 Second matrix.
* \return The resulting matrix.
* \exception panic If the sizes of the matrices do not allow multiplication.
 */
func multiplyMatrix(matrix1, matrix2 [][]int) [][]int {
	if len(matrix1[0]) != len(matrix2) {
		panic("Cannot multiply matrices with incompatible dimensions")
	}

	result := make([][]int, len(matrix1))
	for i := range result {
		result[i] = make([]int, len(matrix2[0]))
	}

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix1[0]); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

/*!
* \brief Sums all elements of a matrix.
*
* \param matrix Matrix.
* \return The sum of all elements of the matrix.
 */
func sumMatrixElements(matrix [][]int) int {
	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
	}
	return sum
}

/*!
* \brief Checks if a matrix is symmetric.
*
* \param matrix Matrix.
* \return true if the matrix is symmetric, otherwise false.
 */
func isSymmetricMatrix(matrix [][]int) bool {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}
