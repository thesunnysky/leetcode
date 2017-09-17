package coin

import (
	"testing"
	"fmt"
)

func TestMinCoin(t *testing.T) {
	target := 11
	cost := make([]int, 11+1)
	//value := make([]int, 3)
	num := MinCoin(target, cost)
	fmt.Println("num:", num)
	for k, v := range cost {
		fmt.Println("target:", k, "cost:", v)
	}
}
