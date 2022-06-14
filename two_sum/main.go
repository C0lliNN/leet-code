package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9

	nums2 := []int{3, 2, 4}
	target2 := 6

	nums3 := []int{3, 3}
	target3 := 6

	fmt.Println(twoSum1(nums1, target1))
	fmt.Println(twoSum1(nums2, target2))
	fmt.Println(twoSum1(nums3, target3))
	fmt.Println("=====================")
	fmt.Println(twoSum2(nums1, target1))
	fmt.Println(twoSum2(nums2, target2))
	fmt.Println(twoSum2(nums3, target3))
	fmt.Println("=====================")
	fmt.Println(twoSum3(nums1, target1))
	fmt.Println(twoSum3(nums2, target2))
	fmt.Println(twoSum3(nums3, target3))
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	numsCopy := make([]int, len(nums))
	for i := range nums {
		numsCopy[i] = nums[i]
	}
	sort.Ints(numsCopy)

	start := 0
	end := len(numsCopy) - 1
	for start < end {
		sum := numsCopy[start] + numsCopy[end]
		if sum == target {
			index1, index2 := 0, 0
			for i := range nums {
				if nums[i] == numsCopy[start] && index1 == 0 {
					index1 = i
					continue
				}

				if nums[i] == numsCopy[end] && index2 == 0 {
					index2 = i
				}
			}

			if index1 < index2 {
				return []int{index1, index2}
			} else {
				return []int{index2, index1}
			}
		}

		if sum < target {
			start++
		} else {
			end--
		}
	}

	return nil
}

func twoSum3(nums []int, target int) []int {
	numbersByIndex := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if index1, ok := numbersByIndex[target-nums[i]]; ok {
			return []int{index1, i}
		}
		numbersByIndex[nums[i]] = i
	}

	return nil
}
