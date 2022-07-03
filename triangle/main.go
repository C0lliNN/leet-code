package main

func main() {

}

func minimumTotal(triangle [][]int) int {
	memo := make([][]int, len(triangle))
	for i := range memo {
		memo[i] = make([]int, i+1)
	}

	return helper(triangle, 0, 0, memo)
}

func helper(triangle [][]int, row, column int, memo [][]int) int {
	if row >= len(triangle) || column >= len(triangle) {
		return 0
	}

	if memo[row][column] > 0 {
		return memo[row][column]
	}

	current := triangle[row][column]

	option1 := current + helper(triangle, row+1, column, memo)
	option2 := current + helper(triangle, row+1, column+1, memo)

	min := option1
	if option2 < option1 {
		min = option2
	}

	memo[row][column] = min

	return min
}

func minimumTotal2(triangle [][]int) int {
	// We can take a Greedy approach starting at the bottom and storing it on the next row up

	// Note: r starts a len-2 because the last row is len-1 and we'll be looking down one row!
	for r := len(triangle) - 2; r >= 0; r-- {
		for c := range triangle[r] {
			// figure out what's the min sum for this node by looking at the possible 2 spots in the next row
			// ie: get the greedy min next step on the next row and add it to the current [r][c] in the triangle
			triangle[r][c] += min(triangle[r+1][c], triangle[r+1][c+1])
		}
	}
	// at this point, the min sum should be at the top
	return triangle[0][0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
