package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob2([]int{1, 2, 3, 1}))
	fmt.Println(rob2([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	return robHelper(nums, len(nums)-1, make([]int, len(nums)))
}

func robHelper(nums []int, i int, memo []int) int {
	if i < 0 {
		return 0
	}

	if memo[i] != 0 {
		return memo[i]
	}

	loot1 := robHelper(nums, i-2, memo) + nums[i]
	loot2 := robHelper(nums, i-1, memo)

	var max int
	if loot1 >= loot2 {
		max = loot1
	} else {
		max = loot2
	}

	memo[i] = max
	return max
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	memo := make([]int, len(nums)+1)
	memo[0] = 0
	memo[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		val := nums[i]

		max := memo[i]
		if memo[i-1]+val > max {
			max = memo[i-1] + val
		}

		memo[i+1] = max
	}

	return memo[len(nums)]
}
