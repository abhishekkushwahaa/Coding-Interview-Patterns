package main

import "fmt"

func PrefixSum(arr []int) []int {
	n := len(arr)
	prefix := make([]int, n)
	prefix[0] = arr[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}
	return prefix
}

func RangeSum(prefix []int, L, R int) int {
	if L == 0 {
		return prefix[R]
	}
	return prefix[R] - prefix[L-1]
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	prefix := PrefixSum(arr)
	L, R := 1, 3
	result := RangeSum(prefix, L, R)
	fmt.Printf("Sum between index %d and %d: %d\n", L, R, result)
}
