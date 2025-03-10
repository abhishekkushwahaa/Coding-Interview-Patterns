package main

import "fmt"

// TC - O(n) && SC - O(1) => (at most 26 characters)
func numSplits(s string) int {
	leftMap := make(map[byte]int)
	rightMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		rightMap[s[i]]++
	}

	count := 0

	for i := 0; i < len(s)-1; i++ {
		leftMap[s[i]]++
		rightMap[s[i]]--

		// If a character count in rightMap becomes zero, delete it
		if rightMap[s[i]] == 0 {
			delete(rightMap, s[i])
		}

		// Compare the number of unique characters in both parts
		if len(leftMap) == len(rightMap) {
			count++
		}
	}

	return count
}

func main() {
	s := "aacaba"
	fmt.Println(numSplits(s)) // Output: 2
}
