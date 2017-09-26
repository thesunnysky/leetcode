package _7_PermutationsII

import (
	"testing"
	"fmt"
)

var input = [][]int{
	{1, 1, 2},
}

func TestPremutationII(t *testing.T) {
	for _, v := range input {
		res := permuteUnique(v)
		fmt.Println(res)
	}
}
