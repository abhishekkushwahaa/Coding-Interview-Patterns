package main

import "fmt"

func permute(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*result = append(*result, permutation)
		return
	}

	seen := make(map[int]bool) // To avoid redundant swaps
	for i := start; i < len(nums); i++ {
		if seen[nums[i]] { // Skip duplicate numbers
			continue
		}
		seen[nums[i]] = true

		nums[start], nums[i] = nums[i], nums[start] // Swap
		permute(nums, start+1, result)              // Recursion
		nums[start], nums[i] = nums[i], nums[start] // Backtracking
	}
}

func getPermutations(nums []int) [][]int {
	var result [][]int
	permute(nums, 0, &result)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	permutations := getPermutations(nums)
	for _, permutation := range permutations {
		fmt.Println(permutation)
	}
}

// TC - O(n!)
// SC - O(n)
