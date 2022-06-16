package main

import "fmt"

func main() {
	numRows1 := 5
	numRows2 := 1

	fmt.Println(generate(numRows1))
	fmt.Println(generate(numRows2))

	fmt.Println(generateRecursive(numRows1))
	fmt.Println(generateRecursive(numRows2))
}

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}

	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 || i == j {
				triangle[i][j] = 1
				continue
			}

			triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
		}
	}

	return triangle
	//      11
	//    21 22
	//   31 32 33
	//  41 42 43 44
	// 51 52 53 54 55
}

func generateRecursive(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1}}
	}

	newRow := make([]int, numRows)
	triangle := append(generateRecursive(numRows-1), newRow)

	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			newRow[i] = 1
			continue
		}

		newRow[i] = triangle[numRows-1][i] + triangle[numRows-1][i-1]
	}

	return triangle
}
