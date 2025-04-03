package main

import "fmt"

// TC - O(log n) && SC - O(1)
func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	first := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			first = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return first
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	last := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			last = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return last
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println("First and Last Occurrences:", searchRange(nums, target)) // Output: [3, 4]
}
