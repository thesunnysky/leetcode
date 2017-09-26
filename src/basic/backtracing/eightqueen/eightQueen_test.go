package eightqueen

import (
	"testing"
	"fmt"
)

func TestEightQueen(t *testing.T) {
	num := eightQueen(8)
	fmt.Println(num)
}
