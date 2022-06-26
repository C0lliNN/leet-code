package main

import (
	"fmt"
)

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println("*********************")
	fmt.Println(lengthOfLongestSubstring2(s1))
	fmt.Println(lengthOfLongestSubstring2(s2))
	fmt.Println(lengthOfLongestSubstring2(s3))
	fmt.Println(lengthOfLongestSubstring2("a"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	longestSubstring := make(map[rune]bool)

	for i := 0; i < len(s); i++ {
		currentSubstring := make(map[rune]bool)
		for j := i; j < len(s); j++ {
			if currentSubstring[rune(s[j])] {
				break
			}

			currentSubstring[rune(s[j])] = true
		}

		if len(currentSubstring) > len(longestSubstring) {
			longestSubstring = currentSubstring
		}
	}

	return len(longestSubstring)
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	longestSubstring, currentSubstring := make(map[rune]bool), make(map[rune]bool)

	for i := 0; i < len(s); i++ {
		if currentSubstring[rune(s[i])] {
			if len(currentSubstring) > len(longestSubstring) {
				longestSubstring = currentSubstring
			}
			currentSubstring = make(map[rune]bool)
		}

		currentSubstring[rune(s[i])] = true
	}

	return len(longestSubstring)
}