package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("adc", "dcda"))
	fmt.Println(checkInclusion("hello", "oolleoooleh"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	for i, j := 0, len(s1); j <= len(s2); i, j = i + 1, j + 1 {
		chars := make(map[rune]int)

		for _, char := range s2[i:j] {
			chars[char]++
		}

		found := true
		for _, char := range s1 {
			chars[char]--
			if chars[char] < 0 {
				found = false
				break
			}
		}

		if found {
			return true
		}
	}

	return false
}
