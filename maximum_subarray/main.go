package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums2 := []int{1}
	nums3 := []int{5, 4, -1, 7, 8}

	fmt.Println(maxSubArray1(nums1))
	fmt.Println(maxSubArray1(nums2))
	fmt.Println(maxSubArray1(nums3))

	fmt.Println(maxSubArray2(nums1))
	fmt.Println(maxSubArray2(nums2))
	fmt.Println(maxSubArray2(nums3))

	fmt.Println(maxSubArray3(nums1))
	fmt.Println(maxSubArray3(nums2))
	fmt.Println(maxSubArray3(nums3))
}

func maxSubArray1(nums []int) int {
	maxSum := math.MinInt

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func maxSubArray2(nums []int) int {
	maxSum := math.MinInt

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func maxSubArray3(nums []int) int {
	best := math.MinInt
	sum := 0

	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > best {
			best = sum
		}
	}

	return best
}
