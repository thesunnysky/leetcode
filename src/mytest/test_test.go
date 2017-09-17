package mytest

import (
	"testing"
	"fmt"
)

func Test_test(t *testing.T) {
	Test()
}

func TestAppend(t *testing.T) {
	a := make([]int, 10, 20)
	for i := 0; i < 1; i++ {
		myAppend(a)
	}
	fmt.Println(a)
}
