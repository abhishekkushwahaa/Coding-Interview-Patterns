package main

import "fmt"

func canCut(wood []int, H int, k int) bool {
	total := 0
	for _, w := range wood {
		if w > H {
			total += w - H
		}
	}
	return total >= k
}

// TC - O(n log max[wood]) && SC - O(1)
func maxCutHeight(wood []int, k int) int {
	left, right := 0, 0
	for _, w := range wood {
		if w > right {
			right = w
		}
	}

	maxH := 0
	for left <= right {
		mid := left + (right-left)/2
		if canCut(wood, mid, k) {
			maxH = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return maxH
}

func main() {
	wood := []int{4, 42, 40, 26, 46}
	k := 20
	fmt.Println("Max Cut Height:", maxCutHeight(wood, k)) // Output: 36
}
