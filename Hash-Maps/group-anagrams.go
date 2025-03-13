package main

import (
	"fmt"
	"strings"
)

// TC - O(n*k) && SC - O(n*k)
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		// Create a frequency count of characters (26 letters for lowercase a-z)
		count := make([]int, 26)

		// Count the frequency of each character
		for _, ch := range str {
			count[ch-'a']++
		}

		// Convert count to a string key (e.g., "1#0#0#2#..." to represent character frequencies)
		key := ""
		for _, c := range count {
			key += fmt.Sprintf("%d#", c)
		}

		// Append the original string to its group
		anagramMap[key] = append(anagramMap[key], str)
	}

	// Collect the result from the map
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	for _, group := range result {
		fmt.Println(strings.Join(group, ", "))
	}
}
