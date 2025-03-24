package main

import "fmt"

func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// TC - O(logn) && SC - O(1)
func isHappy(n int) bool {
	slow, fast := n, getNext(n)

	// Detect cycle using two pointers
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	// If fast reaches 1, it's a happy number
	return fast == 1
}

func main() {
	fmt.Println(isHappy(19)) // Output: true
	fmt.Println(isHappy(2))  // Output: false
}
