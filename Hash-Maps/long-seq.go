package main

import "fmt"

// TC - O(n) && SC - O(n)

func longestConsecutiveSeq(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)

	// Step 1: Insert all numbers into a set
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0

	// Step 2: Find the longest consecutive sequence
	for num := range numSet {

		// Start a sequence only if num-1 is NOT in the set (ensures it's a new sequence)
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			// Count consecutive numbers
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			// Update longest sequence
			if currentLength > longest {
				longest = currentLength
			}
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutiveSeq(nums)) // Output: 4
}
