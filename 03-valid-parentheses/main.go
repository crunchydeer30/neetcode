package main

import "fmt"

func main() {
	fmt.Println("()", isValid("()"))
	fmt.Println("()[]{}", isValid("()[]{}"))
	fmt.Println("(]", isValid("(]"))
	fmt.Println("([])", isValid("([])"))
	fmt.Println("([)]", isValid("([)]"))
	fmt.Println("()", isValid("()"))
}

func isValid(s string) bool {
	stack := []rune{}
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, parentheses := range s {
		if match, found := hash[parentheses]; found {
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, parentheses)
		}
	}

	return len(stack) == 0
}
