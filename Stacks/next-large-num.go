package main

import "fmt"

// TC - O(n) && SC - O(n)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = nums[stack[len(stack)-1]]
		}

		stack = append(stack, i)
	}

	return result
}

func main() {
	nums := []int{4, 5, 2, 10, 8}
	fmt.Println(nextGreaterElements(nums)) // Output: [5 10 10 -1 -1]
}
