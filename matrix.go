package matrix

import (
	"math/rand"
	"time"
)

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
