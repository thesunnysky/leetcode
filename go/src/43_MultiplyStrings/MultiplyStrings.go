package MultiplyStrings

import (
	"strings"
	"fmt"
	"strconv"
	"bytes"
)


func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return "0"
	}

	if len(num1) > len(num2) {
		num2 = strings.Join([]string{"0"}, num2)
		fmt.Println(num2)
	} else if len(num2) > len(num1) {
		num1 = strings.Join([]string{"0"}, num1)
		fmt.Println(num1)
	}

	return doMul(num1, 0, num2, 0);
}

func doMul(num1 string, power1 int, num2 string, power2 int) string {
	if len(num1) == 0 || num1 == "0" {
		return "0"
	}
	if len(num2) == 0 || num2 == "0" {
		return "0"
	}

	if len(num1) == 1 && len(num2) == 1 {
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		temp := a * b
		return genNum(strconv.Itoa(temp), power2)
	} else {

	}

	part1 := doMul(num1[0:1], len(num1)-1, num2[0:1], len(num2)-1)
	part2 := doMul(genNum(num1[0:1], len(num1)-1), 0, num2[1:], 0)
	part3 := doMul(num1[1:], 0, genNum(num2[0:1], len(num1)-1), 0)
	part4 := doMul(num1[1:], 0, num2[1:], 0)
	return Add(Add(part1, part2), Add(part3, part4))
}

func genNum(num string, power int) string {
	var buf bytes.Buffer
	buf.WriteString(num)
	for i := 0; i < power; i++ {
		buf.WriteString("0")
	}
	return buf.String()
}

func Add(numInA string, numInB string) string {
	num1, num2 := Align(numInA, numInB)
	var buf bytes.Buffer
	in := 0
	for i := len(num1) - 1; i >= 0; i-- {
		numA, _ := strconv.Atoi(num1[i-1:i])
		numB, _ := strconv.Atoi(num2[i-1:i])
		sum := numA + numB + in
		if sum >= 10 {
			sum = 1
		} else {
			sum = 0
		}
		buf.WriteString(strconv.Itoa(sum % 10))
	}
	var bufRtn bytes.Buffer
	for i := range buf.String() {
		bufRtn.WriteString(string(i))
	}
	return bufRtn.String()
}

func backup(num1 string, power1 int, num2 string, power2 int) string {
	part1 := doMul(num1[0:1], len(num1)-1, num2[0:1], len(num2)-1)
	part2 := doMul(num1[0:1], len(num1)-1, num2[1:], 0)
	part3 := doMul(num1[1:], 0, num2[0:1], len(num1)-1)
	part4 := doMul(num1[1:], 0, num2[1:], 0)
	return part1 + part2 + part3 + part4
}

func Align(num1 string, num2 string) (string, string) {
	var buf bytes.Buffer
	if len(num1) > len(num2) {
		for i := 0; i < len(num1)-len(num2); i++ {
			buf.WriteString("0")
		}
		buf.WriteString(num2)
		return num1, buf.String()
	} else if len(num1) < len(num2) {
		for i := 0; i < len(num2)-len(num1); i++ {
			buf.WriteString("0")
		}
		return buf.String(), num2
	}
	return num1, num2
}
