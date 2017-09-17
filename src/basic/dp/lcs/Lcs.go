package lcs

import (
	"fmt"
)

func LcsLength(x []rune, y []rune) ([][]int, [][]string) {

	m := len(x)
	n := len(y)

	a := make([][]int, m+1)
	b := make([][]string, m+1)

	//malloc mem and init a
	for i := 0; i < m+1; i ++ {
		for j := 0; j < n+1; j ++ {
			a[i] = append(a[i], 0)
			b[i] = append(b[i], "")
		}
	}
	x = append(make([]rune, 1), x...)
	y = append(make([]rune, 1), y...)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)

	for i := 1; i < len(x); i ++ {
		for j := 1; j < len(y); j++ {
			if x[i] == y[j] {
				a[i][j] = a[i-1][j-1] + 1
				b[i][j] = "diag"
			} else if a[i-1][j] >= a[i][j-1] {
				a[i][j] = a[i-1][j]
				b[i][j] = "upup"
			} else {
				a[i][j] = a[i][j-1]
				b[i][j] = "left"
			}
		}
	}
	return a, b
}

func print2DArrayStr(data [][]string) {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func print2DArrayInt(data [][]int) {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}
