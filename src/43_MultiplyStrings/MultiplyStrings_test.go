package MultiplyStrings

import (
	"testing"
	"fmt"
)

var num1 = []string{"0", "10", "99", "999", "9999"}
var num2 = []string{"10", "10", "99", "999", "9999"}

func TestAdd(t *testing.T) {
	sum := Add("123", "123")
	fmt.Println(sum)
}

func TestMultiply(t *testing.T) {
	result := multiply("123", "456")
	fmt.Println(result)
}

func TestDoMul2(t *testing.T) {
	result := DoMul2("1", 1, "120", 0)
	fmt.Println(result)
}

func TestMulply2(t *testing.T) {
	result := ""
	result = multiply2("0", "0")
	fmt.Println(result)
	result = multiply2("1", "0")
	fmt.Println(result)
	result = multiply2("99", "99")
	fmt.Println(result)
	result = multiply2("99", "999")
	fmt.Println(result)
}

func TestMultiply3(t *testing.T) {
	result := ""
	result = multiply3("99", "9")
	fmt.Println(result)
	//for i := 0; i <= len(num1)-1; i++ {
	//	result = multiply3(num1[i], num2[i])
	//	fmt.Println("num1:", num1[i], "num2", num2[i], "product:", result)
	//}
}
