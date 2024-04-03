package main

func dfs(board [][]byte, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[idx] {
		return false
	}
	tmp := board[i][j]
	board[i][j] = ' ' // mark as visited
	if dfs(board, word, i+1, j, idx+1) || dfs(board, word, i-1, j, idx+1) || dfs(board, word, i, j+1, idx+1) || dfs(board, word, i, j-1, idx+1) {
		return true
	}
	board[i][j] = tmp
	return false
}

func exist(board [][]byte, word string) bool {
	firstCharIIdx := make([]int, 0)
	firstCharJIdx := make([]int, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				firstCharIIdx = append(firstCharIIdx, i)
				firstCharJIdx = append(firstCharJIdx, j)
			}
		}
	}
	for i := 0; i < len(firstCharIIdx); i++ {
		if dfs(board, word, firstCharIIdx[i], firstCharJIdx[i], 0) {
			return true
		}
	}
	return false
}

func main() {
	word := "ABCCED"
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exist(board, word)
}
