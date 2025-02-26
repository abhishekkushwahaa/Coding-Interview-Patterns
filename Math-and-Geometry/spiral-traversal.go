package main

import "fmt"

func spiralTraverse(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var result []int
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for left <= right && top <= bottom {
		// Traverse from left to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse from top to bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse from right to left
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse from bottom to top
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	result := spiralTraverse(matrix)
	fmt.Println("Spiral Order:", result)
}

// TC - O(m*n) (Visits each element once)
// SC - O(1)
