package main

import "fmt"

func main() {
	nums1 := []int{56, 12, 3, 2, 12}
	nums2 := []int{56, 12, 3, 5, 15}

	fmt.Println(containsDuplicate1(nums1))
	fmt.Println(containsDuplicate1(nums2))

	fmt.Println(containsDuplicate2(nums1))
	fmt.Println(containsDuplicate2(nums2))
}

// Time Complexity O (n²)
// Space Complexity O (1)
func containsDuplicate1(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// Time complexity O(n)
// Space Complexity O(n)
func containsDuplicate2(nums []int) bool {
	numbersPresent := make(map[int]bool)

	for _, num := range nums {
		if numbersPresent[num] {
			return true
		}
		numbersPresent[num] = true
	}

	return false
}
