package _9_WrodSearch

func exist(board [][]byte, word string) bool {
	if len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j++ {
			if rune(word[0]) == rune(board[i][j]) {
				board[i][j] = '_'
				if findWord(board, i, j, word, 1) {
					board[i][j] = byte(word[0])
					return true
				}
				board[i][j] = byte(word[0])
			}
		}
	}
	return false
}

func findWords(board [][]byte, words []string) []string {
	keymap := make(map[string]int)
	result := make([]string, 0)
	for _, v := range words {
		if exist(board, v) {
			keymap[v] = 1
		}
	}
	for ret := range keymap {
		result = append(result,ret)
	}
	return result
}

func findWord(board [][]byte, i int, j int, word string, index int) bool {
	if index >= len(word) {
		return true
	}

	rowLimit, colLimit := len(board), len(board[0])
	if i-1 >= 0 && byte(word[index]) == board[i-1][j] {
		board[i-1][j] = '_'
		if findWord(board, i-1, j, word, index+1) {
			board[i-1][j] = byte(word[index])
			return true
		}
		board[i-1][j] = byte(word[index])
	}
	if i+1 < rowLimit && byte(word[index]) == board[i+1][j] {
		board[i+1][j] = '_'
		if findWord(board, i+1, j, word, index+1) {
			board[i+1][j] = byte(word[index])
			return true
		}
		board[i+1][j] = byte(word[index])
	}
	if j-1 >= 0 && byte(word[index]) == board[i][j-1] {
		board[i][j-1] = '_'
		if findWord(board, i, j-1, word, index+1) {
			board[i][j-1] = byte(word[index])
			return true
		}
		board[i][j-1] = byte(word[index])
	}
	if j+1 < colLimit && byte(word[index]) == board[i][j+1] {
		board[i][j+1] = '_'
		if findWord(board, i, j+1, word, index+1) {
			board[i][j+1] = byte(word[index])
			return true
		}
		board[i][j+1] = byte(word[index])
	}
	return false
}

func findWord2(board [][]rune, i int, j int, word string) bool {
	if rune(word[0]) == board[i][j] {
		lastI, lastJ := i, j
		rowLimit, colLimit := len(board), len(board[0])
		index := 1
		for index < len(word) {
			if i-1 >= 0 && rune(word[index]) == board[i-1][j] {
				lastI, i = i, i-1
				index++
			} else if i+1 < rowLimit && rune(word[index]) == board[i+1][j] {
				lastI, i = i, i+1
				index++
			} else if j-1 >= 0 && rune(word[index]) == board[i][j-1] {
				lastJ, j = j, j-1
				index++
			} else if j+1 < colLimit && rune(word[index]) == board[i][j+1] {
				lastJ, j = j, j+1
				index++
			} else {
				i = lastI
				j = lastJ
				index--
			}
			if index < 1 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
