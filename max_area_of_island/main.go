package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	for r := range grid {
		for c := range grid[r] {
			temp := area(grid, r, c)
			if  temp > result {
				result = temp
			}
		}
	}

	return result
}

func area(grid [][]int, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}

	grid[r][c] = 0

	return 1 + area(grid, r+1, c) + area(grid, r-1, c) + area(grid, r, c-1) + area(grid, r, c+1)
}
