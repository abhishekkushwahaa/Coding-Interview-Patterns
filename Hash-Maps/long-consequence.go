package main

import "fmt"

// TC - O(n) && SC - O(n)
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	longest := 0

	for _, num := range nums {
		// Only check if it's the start of a sequence
		if !set[num-1] {
			currentNum := num
			currentStreak := 1

			for set[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}

	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // Output: 4
}
