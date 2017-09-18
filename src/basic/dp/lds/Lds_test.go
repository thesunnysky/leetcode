package lds

import (
	"testing"
	"fmt"
)

var s = []int{1, 2, 3, 9, 27}

func TestLargestDevide(t *testing.T) {
	length := largestDivisibleSubset(s)
	fmt.Println("length:", length)
}
