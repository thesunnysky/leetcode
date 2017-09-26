package _7_SudokuSolver

import (
	"fmt"
)

func solveSudoku(board [][]byte) bool {
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rune(board[i][j]) == '.' {
				candidates := findCandidates(board, i, j)
				for _, v := range candidates {
					board[i][j] = v
					if solveSudoku(board) {
						return true
					} else {
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func findCandidates(board [][]byte, row int, col int) []byte {
	set := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rowLimit, colLimit := len(board), len(board[0])
	for i := 0; i < rowLimit; i++ {
		if rune(board[i][col]) != '.' {
			set[int(board[i][col]-'0')] = 0
		}
	}
	for j := 0; j < colLimit; j++ {
		if rune(board[row][j]) != '.' {
			set[int(board[row][j]-'0')] = 0
		}
	}
	rowLower, rowUpper := (row/3)*3, (row/3)*3+3
	colLower, colUpper := (col/3)*3, (col/3)*3+3
	for i := rowLower; i < rowUpper; i++ {
		for j := colLower; j < colUpper; j++ {
			if rune(board[i][j]) != '.' {
				set[int(board[i][j]-'0')] = 0
			}
		}
	}
	res := make([]byte, 0)
	for _, v := range set {
		if v != 0 {
			res = append(res, byte(v)+'0')
		}
	}
	return res
}

func printBoard(board [][]byte) {
	fmt.Println("=============")
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Print(int(board[i][j]-'0'), " ")
			//fmt.Print(rune(board[i][j]), " ")
		}
		fmt.Println()
	}
}

func toIntArray(data []byte) []int {
	res := make([]int, 0)
	for _, v := range data {
		res = append(res, int(v-'0'))
	}
	return res
}
