package main

import "fmt"

func main() {
	s1 := "leetcode"
	s2 := "loveleetcode"
	s3 := "aabb"

	fmt.Println(firstUniqChar1(s1))
	fmt.Println(firstUniqChar1(s2))
	fmt.Println(firstUniqChar1(s3))
	fmt.Println(firstUniqChar2(s1))
	fmt.Println(firstUniqChar2(s2))
	fmt.Println(firstUniqChar2(s3))
}

func firstUniqChar1(s string) int {
	for i := 0; i < len(s); i++ {
		found := false
		for j := 0; j < len(s); j++ {
			if i != j && s[i] == s[j] {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}

	return -1
}

func firstUniqChar2(s string) int {
	letters := make([]int, 26)

	for i := range s {
		letters[s[i] - 97]++
	}

	for i := range s {
		if letters[s[i] - 97] == 1 {
			return i
		}
	}

	return -1
}

