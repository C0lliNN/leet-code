package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 5, 6}
	target1 := 5

	nums2 := []int{1, 2, 5, 6}
	target2 := 2

	nums3 := []int{1, 3, 5, 6}
	target3 := 7

	fmt.Println(searchInsert(nums1, target1))
	fmt.Println(searchInsert(nums2, target2))
	fmt.Println(searchInsert(nums3, target3))
}

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}
