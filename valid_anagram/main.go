package main

import "fmt"

func main() {
	s1 := "anagram"
	t1 := "nagaram"

	s2 := "rat"
	t2 := "car"

	s3 := "nl"
	t3 := "cx"

	fmt.Println(isAnagram(s1, t1))
	fmt.Println(isAnagram(s2, t2))
	fmt.Println(isAnagram(s3, t3))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lettersInS := make([]int, 26)
	lettersInT := make([]int, 26)
	for i := 0; i < len(s); i++ {
		lettersInS[s[i] - 97]++
		lettersInT[t[i] - 97]++
	}

	for i := 0; i < 26; i++ {
		if lettersInS[i] != lettersInT[i] {
			return false
		}
	}

	return true
}