package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(search(nums, 9))
	fmt.Println(search(nums, 2))
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (end + start) / 2

		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
