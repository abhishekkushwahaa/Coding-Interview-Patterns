package main

import "fmt"

// TC - O(n) && SC - O(k); k - size of the character set
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Map to store the last seen index of each character
	charIndex := make(map[byte]int)
	maxLength, left := 0, 0

	// Sliding window - right pointer expands the window
	for right := 0; right < n; right++ {
		currentChar := s[right]
		if lastSeenIndex, found := charIndex[currentChar]; found && lastSeenIndex >= left {
			left = lastSeenIndex + 1
		}

		// Store/update the index of the current character
		charIndex[currentChar] = right

		// Calculate the length of the current window and update maxLength
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3 ("abc")
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
	fmt.Println(lengthOfLongestSubstring("abcdef"))   // Output: 6 ("abcdef")
}
