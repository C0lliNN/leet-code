package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	matrix2 := [][]int{{1}}
	matrix3 := [][]int{{1, 1}}
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))
	fmt.Println(searchMatrix(matrix2, 2))
	fmt.Println(searchMatrix(matrix3, 2))
}

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}
