package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m1 := 3
	nums2 := []int{2, 5, 6}
	n1 := 3

	nums3 := []int{1}
	m2 := 1
	nums4 := []int{}
	n2 := 0

	nums5 := []int{0}
	m3 := 0
	nums6 := []int{1}
	n3 := 1

	nums7 := []int{4, 0, 0, 0, 0, 0}
	m4 := 1
	nums8 := []int{1, 2, 3, 5, 6}
	n4 := 5

	merge1(nums1, m1, nums2, n1)
	fmt.Println(nums1)
	merge1(nums3, m2, nums4, n2)
	fmt.Println(nums3)
	merge1(nums5, m3, nums6, n3)
	fmt.Println(nums5)
	merge1(nums7, m4, nums8, n4)
	fmt.Println(nums7)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	mergedArray := make([]int, m+n)

	i, j, counter := 0, 0, 0
	for i < m && j < n && counter < m+n {
		if nums1[i] <= nums2[j] {
			mergedArray[counter] = nums1[i]
			i++
		} else {
			mergedArray[counter] = nums2[j]
			j++
		}

		counter++
	}

	for k := i; k < m; k++ {
		mergedArray[counter] = nums1[k]
		counter++
	}

	for k := j; k < n; k++ {
		mergedArray[counter] = nums2[k]
		counter++
	}

	for k := range mergedArray {
		nums1[k] = mergedArray[k]
	}
}
