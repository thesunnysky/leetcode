package lcs

import "testing"

var a = []rune{'A', 'B', 'C', 'B', 'D', 'A', 'B'}
var b = []rune{'B', 'D', 'C', 'A', 'B', 'A'}

func TestLcsLength(t *testing.T) {
	a, b := LcsLength(a, b)
	print2DArrayInt(a)
	print2DArrayStr(b)
}
