package operations

import "fmt"

func printMatrix(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Print(matrix[i][j])
			fmt.Print("\t")
		}
		fmt.Print("\n\n")
	}
}

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
