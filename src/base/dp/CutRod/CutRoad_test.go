package CutRod

import (
	"testing"
	"fmt"
)

var p = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
var len = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestCutRodTopDown(t *testing.T) {
	for _, v := range len {
		max := CutRodTopDown(p, v)
		fmt.Println("length:", v, "value:", max)
	}
}

func TestMemCutRoadTopDown(t *testing.T) {
	mem := make([]int, 11)
	for _, v := range len {
		max := MemCutRoadTopDown(p, v, mem)
		fmt.Println("length:", v, "value:", max)
	}
	for k, v := range mem {
		fmt.Println("key:", k, "value:", v)
	}
}

func TestMemCutRoadDownTop(t *testing.T) {
	mem := make([]int, 11)
	for _, v := range len {
		max := MemCutRoadDownTop(p, v, mem)
		fmt.Println("length:", v, "value:", max)
	}
	for k, v := range mem {
		fmt.Println("key:", k, "value:", v)
	}
}
