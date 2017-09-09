package _43_MultiplyStrings

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	sum := Add("123", "123")
	fmt.Println(sum)
}

func TestMultiply(t *testing.T) {
	result := Multiply("100", "10")
	fmt.Println(result)
}
