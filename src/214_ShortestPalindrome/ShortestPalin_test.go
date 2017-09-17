package ShortestPalindrome

import (
	"fmt"
	"testing"
)

var testCase = []string{"", "a", "ab", "aa", "abcd", "bcba"}

func TestMysol(t *testing.T) {
	for _, str := range testCase {
		result := shortestPalindrome(str)
		fmt.Println("src:", str, "result:", result)
	}
}

func TestMysolSingle(t *testing.T) {
	str := "abbacd"
	result := shortestPalindrome(str)
	fmt.Println("src:", str, "result:", result)
}

func TestMysolSample(t *testing.T) {
	str := "abbacd"
	result := shortestPalindromeSample(str)
	fmt.Println("src:", str, "result:", result)
}
