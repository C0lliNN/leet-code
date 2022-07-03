package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	fmt.Println(permutations)
}

func permute(nums []int) [][]int {
	permutations := make([][]int, 0)
	perm := make([]int, 0)

	backtracking(&permutations, &perm, nums, make([]bool, len(nums)))
	return permutations
}

func backtracking(permutations *[][]int, perm *[]int, nums []int, chosen []bool) {
	if len(*perm) == len(nums) {
		newPerm := make([]int, len(*perm))
		for i := range *perm {
			newPerm[i] = (*perm)[i]
		}
		*permutations = append(*permutations, newPerm)
		return
	}

	for i := range nums {
		if chosen[i] {
			continue
		}

		*perm = append(*perm, nums[i])
		chosen[i] = true

		backtracking(permutations, perm, nums, chosen)

		*perm = (*perm)[:len(*perm)-1]
		chosen[i] = false

	}
}
