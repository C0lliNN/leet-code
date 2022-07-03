package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(32))
	fmt.Println(isPowerOfTwo(33))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n & (n - 1) == 0
}