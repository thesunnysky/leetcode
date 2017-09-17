package mytest

import "fmt"

func Test() {
	fmt.Println("Test")
}

func myAppend(a []int) []int {
	//fmt.Println("before Addr:", &a)
	a = append(a, 1)
	//fmt.Println("after Addr:", &a)
	//a[0] = 1
	return a
}
