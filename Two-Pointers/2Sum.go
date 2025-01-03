// Source: https://leetcode.com/problems/two-sum/
package Two_Pointers

import "sort"

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Time Complexity - O(n^2)

func twoSumOpt(nums []int, target int) []int {
	numsWithIndex := make([][2]int, len(nums))
	for i, num := range nums {
		numsWithIndex[i] = [2]int{num, i}
	}

	sort.Slice(numsWithIndex, func(i, j int) bool {
		return numsWithIndex[i][0] < numsWithIndex[j][0]
	})

	left, right := 0, len(nums)-1
	for left < right {
		sum := numsWithIndex[left][0] + numsWithIndex[right][0]
		if sum == target {
			return []int{numsWithIndex[left][1], numsWithIndex[right][1]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

// Time Complexity - O(nlogn)
