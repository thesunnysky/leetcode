package lis

import (
	"testing"
	"fmt"
)

//var s = []int{1, -1, 0, 2, -3, 4, -6, -5, 5, 6, -7}
//var s = []int{10, 9, 2, 5, 3, 7, 101, 18}
//var s = []int{10, 9, 2, 5, 3, 4}
var s = []int{1, 3, 6, 7, 9, 4, 10, 5, 6}

func TestLis(t *testing.T) {
	maxLen := lengthOfLIS(s)
	fmt.Println(maxLen, "[]:")
}
