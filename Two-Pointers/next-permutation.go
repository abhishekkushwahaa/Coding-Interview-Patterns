package main

import "fmt"

// TC - O(n) && SC - O(1)
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// Step 1: Find the first decreasing element from the right
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// Step 2: Find the next larger element than nums[i] from the right
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}

		// Step 3: Swap nums[i] and nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 4: Reverse the suffix (from i+1 to end)
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1, 3, 2]

	nums = []int{5, 4, 7, 5, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [5, 5, 2, 3, 4, 7]
}
