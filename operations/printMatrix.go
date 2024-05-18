/*!
* \file printMatrix.go
* \brief Functions for printing matrices to the screen.
 */
package operations

import "fmt"

/*!
* \brief Prints the matrix to the screen.
*
* \param matrix Matrix to be printed.
 */
func printMatrix(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Print(matrix[i][j])
			fmt.Print("\t")
		}
		fmt.Print("\n\n")
	}
}

/*!
* \brief Prints the specified row of the matrix to the screen.
*
* \param matrix Matrix to be printed.
* \param rowIndex Index of the row to be printed.
 */
func printRow(matrix [][]int, rowIndex int) {
	if rowIndex < 0 || rowIndex >= len(matrix) {
		fmt.Println("Invalid row index")
		return
	}
	for _, value := range matrix[rowIndex] {
		fmt.Print(value, "\t")
	}
	fmt.Println()
}
