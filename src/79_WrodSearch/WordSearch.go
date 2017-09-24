package _9_WrodSearch

func wordSearch(board [][]rune, word string) bool {
	if len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j++ {
			if rune(word[0]) == board[i][j] && findWord(board, i, j, word, 1) {
				return true
			}
		}
	}
	return false
}

func findWord(board [][]rune, i int, j int, word string, index int) bool {
	if index > len(word) {
		return true
	}

	rowLimit, colLimit := len(board), len(board[0])
	if i-1 >= 0 && rune(word[index]) == board[i-1][j] {
		if findWord(board, i-1, j, word, index+1) {
			return true
		}
	} else if i+1 < rowLimit && rune(word[index]) == board[i+1][j] {
		if findWord(board, i+1, j, word, index+1) {
			return true
		}
	} else if j-1 >= 0 && rune(word[index]) == board[i][j-1] {
		if findWord(board, i, j-1, word, index+1) {
			return true
		}
	} else if j+1 < colLimit && rune(word[index]) == board[i][j+1] {
		if findWord(board, i, j+1, word, index+1) {
			return true
		}
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
