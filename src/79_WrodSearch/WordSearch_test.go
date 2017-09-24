package _9_WrodSearch

import (
	"testing"
	"fmt"
)

var board = [][]rune{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'C', 'S'},
	{'A', 'D', 'E', 'E'},
}

var words = []string{"SEE", "ABCCED", "SEE", "ABCB"}

func TestWordSearch(t *testing.T) {
	for _, v := range words {
		if wordSearch(board, v) {
			fmt.Println(v, "in board")
		}
	}
}
