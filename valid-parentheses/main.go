package main

import (
	"fmt"
)

func isValid(s string) bool {
	// Create a stack to keep track of opening brackets
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	// The string is valid if the stack is empty
	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(isValid("()"))     // Output: true
	fmt.Println(isValid("()[]{}")) // Output: true
	fmt.Println(isValid("(]"))     // Output: false
	fmt.Println(isValid("([)]"))   // Output: false
	fmt.Println(isValid("{[]}"))   // Output: true
}
