/*!
* \file matrix.go
* \brief Functions for creating matrices.
 */
package matrix

import (
	"math/rand"
	"time"
)

/*!
* \brief Creates a matrix of the specified size with random elements.
*
* \param size Size of the matrix (number of rows and columns).
* \return The generated matrix.
 */
func createMatrix(size int) [][]int {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]int, size)
	for i := range matrix {
		for j := 0; j < size; j++ {
			randomNumber := rand.Intn(100) + 1
			matrix[i] = append(matrix[i], randomNumber)
		}
	}
	return matrix
}

/*!
* \brief Creates a matrix of the specified size with random elements limited by a maximum value.
*
* \param size Size of the matrix (number of rows and columns).
* \param maxValue Maximum value of the matrix element.
* \return The generated matrix.
 */
func createLimitedMatrix(size, maxValue int) [][]int {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]int, size)
	for i := range matrix {
		for j := 0; j < size; j++ {
			randomNumber := rand.Intn(maxValue) + 1
			matrix[i] = append(matrix[i], randomNumber)
		}
	}
	return matrix
}
