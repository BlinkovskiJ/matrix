package operations

func multiplyMatrix(matrix1, matrix2 [][]int) [][]int {
	if len(matrix1[0]) != len(matrix2) {
		panic("Нельзя умножить матрицы разной размерности")
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

func sumMatrixElements(matrix [][]int) int {
	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
	}
	return sum
}

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
