package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	return helper(n, make([]int, n+1))
}

func helper(n int, cache []int) int {
	if n <= 2 {
		return n
	}

	if cache[n] != 0 {
		return cache[n]
	}

	val1 := helper(n-1, cache)
	cache[n-1] = val1

	val2 := helper(n-2, cache)
	cache[n-2] = val2

	return val1 + val2
}
