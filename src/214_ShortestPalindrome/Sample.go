package ShortestPalindrome

import "fmt"

func shortestPalindromeSample(s string) string {
	if len(s) == 0 {
		return ""
	}

	next := make([]int, len(s))
	next[0] = -1

	for i := 0; i < len(s)-1; {
		j := next[i]
		for j >= 0 && s[i] != s[j] {
			fmt.Printf("outer:s[%d]:%c, s[%d]:%c\n", i, s[i], j, s[j])
			j = next[j]
		}
		if j >= 0 {
			fmt.Printf("outer:s[%d]:%c, s[%d]:%c\n", i, s[i], j, s[j])
		}
		i++
		next[i] = j + 1
		fmt.Println("next", next)
	}

	i := 0
	j := len(s) - 1

	for j > i {
		for i >= 0 && s[i] != s[j] {
			i = next[i]
		}
		i++
		j--
	}

	var end int
	if i == j {
		end = i*2 + 1
	} else {
		end = i * 2
	}
	result := make([]byte, 0, 2*len(s)-end)
	for i := len(s) - 1; i >= end; i-- {
		result = append(result, s[i])
	}
	result = append(result, s...)
	return string(result)
}
