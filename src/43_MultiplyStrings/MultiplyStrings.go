package MultiplyStrings

import (
	"strconv"
	"bytes"
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if IsAllZero(num1) || IsAllZero(num2) {
		return "0"
	}
	num1, num2 = Align(num1, num2)

	return DoMul(num1, 0, num2, 0);
}

func DoMul(num1 string, power1 int, num2 string, power2 int) string {
	fmt.Println("num1:", num1, "num2:", num2)
	if IsAllZero(num1) || IsAllZero(num2) {
		return "0"
	}

	num1, num2 = Align(num1, num2)

	if len(num1) == 1 && len(num2) == 1 {
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		temp := a * b
		return genNum(strconv.Itoa(temp), power1+power2)
	} else {
		part1 := DoMul(num1[0:1], len(num1)-1, num2[0:1], len(num2)-1)
		//part2 := DoMul(genNum(num1[0:1], len(num1)-1), 0, num2[1:], 0)
		part2 := DoMul2(num1[0:1], len(num1)-1, num2[1:], 0)
		//part3 := DoMul(num1[1:], 0, genNum(num2[0:1], len(num2)-1), 0)
		part3 := DoMul2(num2[0:1], len(num2)-1, num1[1:], 0)
		part4 := DoMul(num1[1:], 0, num2[1:], 0)
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

func DoMul2(num1 string, power int, num2 string, power2 int) string {
	if IsAllZero(num1) || IsAllZero(num2) {
		return "0"
	}
	var buf bytes.Buffer
	in := 0
	for i := len(num2); i > 0; i-- {
		numA, _ := strconv.Atoi(num1)
		numB, _ := strconv.Atoi(num2[i-1:i])
		product := numA*numB + in
		in = product / 10
		stay := product % 10
		buf.WriteString(strconv.Itoa(stay))
	}
	if in != 0 {
		buf.WriteString(strconv.Itoa(in))
	}
	result := ReverseString(buf.String())
	buf.Reset()
	buf.WriteString(result)
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
		in = sum / 10
		buf.WriteString(strconv.Itoa(sum % 10))
	}
	if in != 0 {
		buf.WriteString(strconv.Itoa(in))
	}
	return removeHeadZero(ReverseString(buf.String()))
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
		buf.WriteString(num1)
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

func IsAllZero(num string) bool {
	for _, ch := range num {
		if !strings.EqualFold(string(ch), "0") {
			return false
		}
	}
	return true
}

func removeHeadZero(num string) string {
	i := 0
	for k, ch := range num {
		if !strings.EqualFold(string(ch), "0") {
			i = k
			break
		}
	}
	return num[i:]
}
