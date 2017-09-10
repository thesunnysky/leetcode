package ShortestPalindrome

import (
	"fmt"
	"bytes"
)

func shortestPalindrome(s string) string {
	src := []rune(s)
	fmt.Println(src)
	from := 0
	for i := len(src) - 1; i >= 0; i-- {
		if isOddCenterOK(src, i) {
			from = i*2 + 1
			break
		} else if isEvenCenterOK(src, i) {
			from = i * 2
			break
		}
	}
	cpy := copy(src, from)
	return cpy + s

}

func copy(src []rune, from int) string {
	var buf bytes.Buffer
	for i := len(src) - 1; i >= from; i-- {
		buf.WriteRune(src[i])
	}
	return buf.String()
}

func isOddCenterOK(src []rune, center int) bool {
	if center >= len(src) || center < 0 {
		return false
	}

	if center == 0 {
		return true
	}

	for i, k := center-1, center+1; i >= 0; i-- {
		if k >= len(src) {
			return false
		}
		if src[i] != src[k] {
			return false
		}
		k++
	}
	return true
}

func isEvenCenterOK(src []rune, center int) bool {
	if center >= len(src) || center < 0 {
		return false
	}

	if center == 0 {
		return false
	}

	for i, k := center-1, center; i >= 0; i-- {
		if k >= len(src) {
			return false
		}
		if src[i] != src[k] {
			return false
		}
		k++
	}
	return true
}
