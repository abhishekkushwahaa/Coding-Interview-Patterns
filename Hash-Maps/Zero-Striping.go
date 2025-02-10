package main

import "fmt"

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])

	// // Using bool Array
	// row := make([]bool, n)
	// col := make([]bool, m)

	// Using Hash - Map
	row := make(map[int]bool, n)
	col := make(map[int]bool, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Println("Original Matrix:")
	printMatrix(matrix)

	setZeroes(matrix)

	fmt.Println("\nMatrix After Setting Zeroes:")
	printMatrix(matrix)
}
