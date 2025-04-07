package main

import (
	"fmt"
	"unicode"
)

// TC - O(n) && SC - O(n)
func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := '+'

	for i, ch := range s {
		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		}

		// If it's an operator or end of string
		if (!unicode.IsDigit(ch) && ch != ' ') || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				last := stack[len(stack)-1]
				stack[len(stack)-1] = last * num
			case '/':
				last := stack[len(stack)-1]
				stack[len(stack)-1] = last / num
			}
			sign = ch
			num = 0
		}
	}

	// Sum up the stack
	result := 0
	for _, val := range stack {
		result += val
	}
	return result
}

func main() {
	expr := "3+2*2"
	fmt.Println("Result:", calculate(expr)) // Output: 7
}
