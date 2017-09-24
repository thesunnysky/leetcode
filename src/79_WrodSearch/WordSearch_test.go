package _9_WrodSearch

import (
	"testing"
	"fmt"
	"bytes"
	"encoding/gob"
)

var board = [][]rune{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'C', 'S'},
	{'A', 'D', 'E', 'E'},
}

func genBoard() [][]byte {
	var board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	return board
}

var words = []string{"ABCCED", "SEE", "ABCB"}

func TestWordSearch(t *testing.T) {
	temp := genBoard()
	for _, v := range words {
		fmt.Println(temp)
		if exist(temp, v) {
			fmt.Println(v, "in board")
		}
	}
}

func TestFindWords(t *testing.T) {
	words := []string{"oath", "pea", "eat", "rain"}
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	result := findWords(board, words)
	fmt.Println(result)
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
