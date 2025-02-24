package main

import "fmt"

func main() {
	// Define matrices
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix2 := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	// Initialize result matrix
	result := make([][]int, len(matrix1))
	for i := range result {
		result[i] = make([]int, len(matrix2[0]))
	}

	// Matrix multiplication
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	// Display the result
	fmt.Println("Resulting matrix:")
	for _, row := range result {
		fmt.Println(row)
	}
}
