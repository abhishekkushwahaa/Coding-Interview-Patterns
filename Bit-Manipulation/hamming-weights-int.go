package main

import "fmt"

// Count bits using Brian Kernighan’s Algorithm - Bitwise Manipulation TC-O(nlogn) & SC-O(1)
func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		n &= (n - 1)
		count++
	}
	return count
}

// Hamming weight for all numbers up to n using DP TC-O(n) & SC-O(n)
func hammingWeightDP(n int) []int {
	dp := make([]int, n+1) // O(n) - Space
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1) // DP formula
	}
	return dp
}

// Hamming weight up to n using Recusive DP TC-O(nlogn) & SC-O(1)
func hammingW(n int) int {
	if n == 0 {
		return 0
	}
	return hammingW(n>>1) + (n & 1)
}

func main() {
	nums := []int{5, 7, 15, 32}

	fmt.Println("Using Bitwise Manipulation:")
	for _, num := range nums {
		fmt.Printf("Hamming Weight of %d: %d\n", num, hammingWeight(num))
	}

	fmt.Println("\nUsing DP:")
	n := 32
	hammingWeights := hammingWeightDP(n)
	for _, num := range nums {
		fmt.Printf("Hamming Weight of %d: %d\n", num, hammingWeights[num])
	}

	fmt.Println("\nUsing Recursive DP:")
	for _, num := range nums {
		fmt.Printf("Hamming Weight of %d: %d\n", num, hammingW(num))
	}
}
