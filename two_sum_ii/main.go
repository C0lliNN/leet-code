package main

import "fmt"

func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{2, 3, 4}

	fmt.Println(twoSum(nums1, 9))
	fmt.Println(twoSum(nums2, 6))
}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers) - 1
	for i < j {
		sum :=  numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}
