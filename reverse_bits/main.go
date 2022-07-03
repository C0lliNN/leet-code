package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseBits(4294967293)) // 964176192
}

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return 0
	}

	result := 0
	for i := 0; i < 32; i++ {
		result <<= 1
		if (num & 1) == 1 {
			result++
		}
		num >>= 1
	}

	return uint32(result)
}
