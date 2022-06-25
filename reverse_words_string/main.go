package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Let's take LeetCode contest"
	s2 := "God Ding"

	fmt.Println([]byte("a b"))

	fmt.Println(reverseWords(s1))
	fmt.Println(reverseWords(s2))

	fmt.Println(reverseWords2(s1))
	fmt.Println(reverseWords2(s2))
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	bytes := make([][]byte, 0)
	for _, word := range words {
		wordBytes := []byte(word)
		for i, j := 0, len(wordBytes) - 1; i < j; i, j = i + 1, j - 1 {
			wordBytes[i], wordBytes[j] = wordBytes[j], wordBytes[i]
		}

		bytes = append(bytes, wordBytes)
	}

	for i := range words {
		words[i] = string(bytes[i])
	}

	return strings.Join(words, " ")
}

func reverseWords2(s string) string {
	bytes := []byte(s)

	i, j := 0, len(bytes) - 1
	for i < j {
		if bytes[i] == 32 {
			i++
			continue
		}
		if bytes[j] == 32 {
			j--
			continue
		}

		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}