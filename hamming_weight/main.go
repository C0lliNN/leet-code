package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(0))
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(32))
	fmt.Println(hammingWeight2(0))
	fmt.Println(hammingWeight2(11))
	fmt.Println(hammingWeight2(32))
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num % 2 == 1 {
			count++
		}
		num /= 2
	}

	return count
}

func hammingWeight2(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num) & 1
		num = num >> uint32(1)
	}

	return count
}

