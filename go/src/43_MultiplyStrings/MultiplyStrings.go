package _43_MultiplyStrings

import (
	"strconv"
	"bytes"
)

func Multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return "0"
	}

	num1, num2 = Align(num1, num2)

	return doMul(num1, 0, num2, 0);
}

func doMul(num1 string, power1 int, num2 string, power2 int) string {
	if len(num1) == 0 || num1 == "0" {
		return "0"
	}
	if len(num2) == 0 || num2 == "0" {
		return "0"
	}
	num1, num2 = Align(num1, num2)

	if len(num1) == 1 && len(num2) == 1 {
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		temp := a * b
		return genNum(strconv.Itoa(temp), power2)
	} else {
		part1 := doMul(num1[0:1], len(num1)-1, num2[0:1], len(num2)-1)
		part2 := doMul(genNum(num1[0:1], len(num1)-1), 0, num2[1:], 0)
		part3 := doMul(num1[1:], 0, genNum(num2[0:1], len(num1)-1), 0)
		part4 := doMul(num1[1:], 0, num2[1:], 0)
		return Add(Add(part1, part2), Add(part3, part4))
	}
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
	for i := len(num1); i > 0; i-- {
		numA, _ := strconv.Atoi(num1[i-1:i])
		numB, _ := strconv.Atoi(num2[i-1:i])
		sum := numA + numB + in
		if sum >= 10 {
			in = 1
		} else {
			in = 0
		}
		buf.WriteString(strconv.Itoa(sum % 10))
	}
	return ReverseString(buf.String())
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

func ReverseString(str string) string {
	runes := []rune(str)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
