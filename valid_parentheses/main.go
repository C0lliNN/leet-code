package main

import "fmt"

func main() {
	s1 := "()"
	s2 := "()[]{}"
	s3 := "(]"

	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
	fmt.Println(isValid(s3))
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastChar := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if char == ')' && lastChar != '(' {
			return false
		}
		if char == ']' && lastChar != '[' {
			return false
		}
		if char == '}' && lastChar != '{' {
			return false
		}
	}

	return len(stack) == 0
}
