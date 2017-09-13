package CutRod

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func printArray(data []int) {
	for k, v := range data {
		fmt.Print(k, ":", v, " ")
	}
	fmt.Println()
}
