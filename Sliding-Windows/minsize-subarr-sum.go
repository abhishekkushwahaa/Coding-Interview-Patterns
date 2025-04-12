package main

import (
	"fmt"
	"math"
)

// TC - O(n) && SC - O(1)
func minSubArrayLen(target int, nums []int) int {
	start := 0
	sum := 0
	minLen := math.MaxInt32

	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		for sum >= target {
			if end-start+1 < minLen {
				minLen = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums)) // Output: 2
}
