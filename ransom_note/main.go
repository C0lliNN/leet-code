package main

import "fmt"

func main() {
	ransomNote1 := "a"
	magazine1 := "b"
	ransomNote2 := "aa"
	magazine2 := "ab"
	ransomNote3 := "aa"
	magazine3 := "aab"

	fmt.Println(canConstruct1(ransomNote1, magazine1))
	fmt.Println(canConstruct1(ransomNote2, magazine2))
	fmt.Println(canConstruct1(ransomNote3, magazine3))
	fmt.Println(canConstruct2(ransomNote1, magazine1))
	fmt.Println(canConstruct2(ransomNote2, magazine2))
	fmt.Println(canConstruct2(ransomNote3, magazine3))
}

func canConstruct1(ransomNote string, magazine string) bool {
	for i := 0; i < len(ransomNote); i++ {
		found := false
		for j := 0; j < len(magazine); j++ {
			if ransomNote[i] == magazine[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	letters := make([]int, 26)

	for i := range magazine {
		letters[magazine[i] - 97]++
	}

	for i := range ransomNote {
		if letters[ransomNote[i] - 97] <= 0 {
			return false
		}
		letters[ransomNote[i] - 97]--
	}

	return true
}