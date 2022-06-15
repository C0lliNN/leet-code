package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersect1(nums1, nums2))
	fmt.Println(intersect1(nums3, nums4))

}

func intersect1(nums1 []int, nums2 []int) []int {
	occurrences := make(map[int]int)

	for _, num := range nums1 {
		occurrences[num]++
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		if occurrences[num] > 0 {
			result = append(result, num)
			occurrences[num]--
		}
	}

	return result
}
