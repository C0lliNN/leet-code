package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{-4, -1, 0, 3, 10}
	nums2 := []int{-7, -3, 2, 3, 11}

	fmt.Println(sortedSquares1(nums1))
	fmt.Println(sortedSquares1(nums2))
}

func sortedSquares1(nums []int) []int {
	for i := range nums {
		if nums[i] < 0 {
			nums[i] = -nums[i]
		}
	}

	sort.Ints(nums)

	for i := range nums {
		nums[i] *= nums[i]
	}

	return nums
}
