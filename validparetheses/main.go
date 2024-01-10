package main

import (
	"fmt"
)

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range s {
		if opening, isOpening := pairs[char]; isOpening {
			stack = append(stack, opening)
		} else {
			// check if the stack is empty or if the current char is not a matching closing bracket for the current opening bracket on the stack
			if len(stack) == 0 || char != stack[len(stack)-1] {
				return false
			}
			// if it matches then remove the last item from the stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
