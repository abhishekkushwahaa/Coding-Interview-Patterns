package main

import "fmt"

// TC - O(n) && SC - O(1)
func isJumpToEnd(nums []int) bool {
	far := 0
	for i := 0; i < len(nums); i++ {
		if i > far {
			return false
		}
		far = max(far, i+nums[i])
		if far >= len(nums)-1 {
			return true
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println(isJumpToEnd(nums1)) // Output: true

	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println(isJumpToEnd(nums2)) // Output: false
}
