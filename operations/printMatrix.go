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
