package eightqueen

import (
	"fmt"
)

var all = 0

func eightQueen(dimension int) {
	chess := genChess(dimension)
	place(chess, 0)
}

func genChess(dimension int) [][]int {
	chess := make([][]int, dimension)
	for i := range chess {
		chess[i] = make([]int, dimension)
	}
	return chess
}

func place(a [][]int, row int) {
	//fmt.Println("finding row", row, "...")
	if row == len(a)-1 {
		rowData := findCol(a, row)
		for i := 0; i < len(rowData); i++ {
			//set col data and print
			if rowData[i] != 0 {
				a[row][i] = rowData[i]
				print(a)
				//reset the row data
				a[row][i] = 0
			}
		}
		return
	} else {
		rowData := findCol(a, row)
		for i := 0; i < len(rowData); i++ {
			if rowData[i] != 0 {
				a[row][i] = rowData[i]
				place(a, row+1)
				a[row][i] = 0
			}
		}
	}
}

func findCol(a [][]int, row int) []int {
	result := make([]int, len(a))
	for i := 0; i < len(a); i ++ {
		if isColOk(a, row, i) && isDiaOk(a, row, i) {
			result[i] = 1
		}
	}
	return result
}

func isColOk(a [][]int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if a[i][col] != 0 {
			return false
		}
	}
	return true
}

func isDiaOk(a [][]int, row int, col int) bool {
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if a[i][j] != 0 {
			return false
		}
		i--
		j--
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0 && j < len(a); {
		if a[i][j] != 0 {
			return false
		}
		i--
		j++
	}
	return true
}

func print(a [][]int) {
	fmt.Println("=======================", all)
	/*	for row := range a {
			fmt.Println(row)
		}*/
	for i := 0; i < len(a); i ++ {
		fmt.Println(a[i])
	}
	fmt.Println("========================")
	all++
}
