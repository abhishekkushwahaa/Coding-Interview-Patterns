package main

import "fmt"

func ClimbingStairsMemo(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if val, exits := memo[n]; exits {
		return val
	}

	memo[n] = ClimbingStairsMemo(n-1, memo) + ClimbingStairsMemo(n-2, memo)

	return memo[n]
}

// Using Top-Down Approach with Memoization: TC - O(n) && SC - O(n)
func ClimbingStairsTopDown(n int) int {
	memo := make(map[int]int)
	return ClimbingStairsMemo(n, memo)
}

// Using Bottom-Up Approach: TC - O(n) && SC - O(1)
func ClimbingStairsBottomUp(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	prev1, prev2 := 1, 1

	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}

func main() {
	fmt.Println(ClimbingStairsTopDown(5))  // Output: 8
	fmt.Println(ClimbingStairsBottomUp(5)) // Output: 8
}
