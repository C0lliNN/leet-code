package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}

func letterCasePermutation(s string) []string {
	permutations := make([]string, 0)
	backtrack(&permutations, []byte(s), 0)
	return permutations
}

func backtrack(permutations *[]string, chars []byte, i int) {
	if i == len(chars) {
		*permutations = append(*permutations, string(chars))
		return
	}

	char := rune(chars[i])

	if unicode.IsLetter(char) {
		chars[i] = byte(unicode.ToUpper(char))
		backtrack(permutations, chars, i+1)
		chars[i] = byte(unicode.ToLower(char))
		backtrack(permutations, chars, i+1)
	} else {
		backtrack(permutations, chars, i+1)
	}
}
