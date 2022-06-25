package main

import "fmt"

func main() {
	s1 := []byte("hello")
	s2 := []byte("Hannah")

	reverseString(s1)
	reverseString(s2)

	fmt.Println(string(s1))
	fmt.Println(string(s2))
}

func reverseString(s []byte)  {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

