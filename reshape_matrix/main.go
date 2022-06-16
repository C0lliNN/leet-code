package main

import "fmt"

func main() {
	mat1 := [][]int{{1, 2}, {3, 4}}
	r1 := 1
	c1 := 4

	mat2 := [][]int{{1, 2}, {3, 4}}
	r2 := 2
	c2 := 4

	fmt.Println(matrixReshape(mat1, r1, c1))
	fmt.Println(matrixReshape(mat2, r2, c2))

	fmt.Println(matrixReshape2(mat1, r1, c1))
	fmt.Println(matrixReshape2(mat2, r2, c2))
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	reshapedMatrixSize := r * c
	currentMatrixSize := len(mat) * len(mat[0])

	if reshapedMatrixSize != currentMatrixSize {
		return mat
	}

	elements := make([]int, 0)
	for _, row := range mat {
		for _, val := range row {
			elements = append(elements, val)
		}
	}

	reshapedMatrix := make([][]int, r)
	for i := range reshapedMatrix {
		reshapedMatrix[i] = make([]int, c)
	}

	counter := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			reshapedMatrix[i][j] = elements[counter]
			counter++
		}
	}

	return reshapedMatrix
}

func matrixReshape2(mat [][]int, r int, c int) [][]int {
	reshapedMatrixSize := r * c
	currentMatrixSize := len(mat) * len(mat[0])

	if reshapedMatrixSize != currentMatrixSize {
		return mat
	}

	reshapedMatrix := make([][]int, r)
	for i := range reshapedMatrix {
		reshapedMatrix[i] = make([]int, c)
	}

	k, l := 0, 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			reshapedMatrix[i][j] = mat[k][l]
			l++
			if l == len(mat[0]) {
				k++
				l = 0
			}
		}
	}

	return reshapedMatrix
}
