package main

import "fmt"

// TC - O(n) && SC - O(1)
func characterReplacement(s string, k int) int {
	// Frequency map for window
	left, maxCount, maxLength := 0, 0, 0

	freq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		maxCount = max(maxCount, freq[s[right]])

		if (right-left+1)-maxCount > k {
			freq[s[left]]--
			left++
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println("Longest length:", characterReplacement(s, k)) // Output: 4
}
