package main

import "fmt"

func insertionIndex(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	target := 6
	index := insertionIndex(arr, target)
	fmt.Printf("Insertion index for %d is %d!\n", target, index)
}
