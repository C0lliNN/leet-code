package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	nums2 := []int{-1, -100, 3, 99}
	k2 := 2

	rotate2(nums1, k1)
	rotate2(nums2, k2)

	fmt.Println(nums1)
	fmt.Println(nums2)
}

func rotate1(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	k %= len(nums)

	for i := 0; i < k; i++ {
		tempArray := make([]int, len(nums))
		for j := range nums {
			tempArray[j] = nums[j]
		}

		for j := 0; j < len(nums); j++ {
			var c int
			if j == len(nums)-1 {
				c = 0
			} else {
				c = j + 1
			}
			nums[c] = tempArray[j]
		}
	}
}

func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
