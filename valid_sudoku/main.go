package main

import "fmt"

func main() {
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board3 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'5', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board4 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '7', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku1(board1))
	fmt.Println(isValidSudoku1(board2))
	fmt.Println(isValidSudoku1(board3))
	fmt.Println(isValidSudoku1(board4))

	fmt.Println("==============")
	fmt.Println(isValidSudoku2(board1))
	fmt.Println(isValidSudoku2(board2))
	fmt.Println(isValidSudoku2(board3))
	fmt.Println(isValidSudoku2(board4))
}

func isValidSudoku1(board [][]byte) bool {
	// check rows
	// check columns
	// check 3 x 3 boxes

	for _, row := range board {
		duplicates := make(map[byte]bool)
		for _, cell := range row {
			if cell == '.' {
				continue
			}
			if duplicates[cell] {
				return false
			}
			duplicates[cell] = true
		}
	}

	for i := 0; i < len(board); i++ {
		duplicates := make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			if board[j][i] == '.' {
				continue
			}
			if duplicates[board[j][i]] {
				return false
			}
			duplicates[board[j][i]] = true
		}
	}

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[i]); j += 3 {
			duplicates := make(map[byte]bool)
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if board[k][l] == '.' {
						continue
					}
					if duplicates[board[k][l]] {
						return false
					}
					duplicates[board[k][l]] = true
				}
			}
		}
	}

	return true
}

// does the same thing as the above method, but using a single loop
func isValidSudoku2(board [][]byte) bool {
	// check rows
	// check columns
	// check 3 x 3 boxes

	columnDuplicates := make(map[int]map[byte]bool)
	subboxesDuplicates := make(map[int]map[byte]bool)
	for i := 0; i < len(board); i++ {
		rowDuplicates := make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			if rowDuplicates[board[i][j]] {
				return false
			}
			rowDuplicates[board[i][j]] = true

			if columnDuplicates[j] == nil {
				columnDuplicates[j] = make(map[byte]bool)
			}

			if columnDuplicates[j][board[i][j]] {
				return false
			}
			columnDuplicates[j][board[i][j]] = true

			subbox := getSubboxFromPosition(i, j)
			if subboxesDuplicates[subbox] == nil {
				subboxesDuplicates[subbox] = make(map[byte]bool)
			}
			if subboxesDuplicates[subbox][board[i][j]] {
				return false
			}
			subboxesDuplicates[subbox][board[i][j]] = true
		}
	}

	return true
}

func getSubboxFromPosition(i, j int) int {
	if i < 3 {
		if j < 3 {
			return 1
		}
		if j < 6 {
			return 2
		}
		return 3
	}
	if i < 6 {
		if j < 3 {
			return 4
		}
		if j < 6 {
			return 5
		}
		return 6
	}
	if j < 3 {
		return 7
	}
	if j < 6 {
		return 8
	}
	return 9
}
