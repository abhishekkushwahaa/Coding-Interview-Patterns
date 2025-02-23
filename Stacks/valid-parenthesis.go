package main

import "fmt"

func validParenthesis(expr string) bool {
	parenthesisMap := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	stacks := []rune{}

	for _, char := range expr {
		if closing, ok := parenthesisMap[char]; ok {
			stacks = append(stacks, closing)
		} else {
			if len(stacks) == 0 || stacks[len(stacks)-1] != char {
				return false
			}
			stacks = stacks[:len(stacks)-1]
		}
	}
	return len(stacks) == 0
}

func main() {
	result := validParenthesis("([]{})")
	fmt.Println(result)
}
