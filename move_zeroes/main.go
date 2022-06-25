package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 0, 3, 12}
	nums2 := []int{0}

	moveZeroes2(nums1)
	moveZeroes2(nums2)

	fmt.Println(nums1)
	fmt.Println(nums2)
}

func moveZeroes(nums []int)  {
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 {
				aux := nums[i]
				nums[i] = nums[j]
				nums[j] = aux
			}
		}
	}
}

func moveZeroes2(nums []int) {
	slow := 0
	for fast := range nums {
		if nums[fast] != 0 && nums[slow] == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}

		if nums[slow] != 0 {
			slow++
		}
	}
}
