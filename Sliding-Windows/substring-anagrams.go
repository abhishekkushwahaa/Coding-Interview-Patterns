package main

import "fmt"

func findAnagrams(s, p string) []int {
	res := []int{}

	if len(s) < len(p) {
		return res
	}

	sDict := make([]int, 26)
	pDict := make([]int, 26)

	for i := 0; i < len(p); i++ {
		pDict[p[i]-'a']++
		sDict[s[i]-'a']++
	}

	start := 0
	if isEqual(pDict, sDict) {
		res = append(res, start)
	}

	for end := len(p); end < len(s); end++ {
		sDict[s[end]-'a']++
		sDict[s[start]-'a']--
		start++
		if isEqual(pDict, sDict) {
			res = append(res, start)
		}
	}

	return res
}

func isEqual(pDist, sDist []int) bool {
	for i := 0; i < 26; i++ {
		if pDist[i] != sDist[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	result := findAnagrams(s, p)
	fmt.Println(result)
}
