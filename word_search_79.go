package dsa

var dir = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var rows = 0
var cols = 0

func isValidPath(x, y int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

func helper(board [][]byte, word string, vis [][]bool, i, x, y int) bool {
	i = i + 1
	if i == len(word) {
		return true
	}
	for _, d := range dir {
		xd := x + d[0]
		yd := y + d[1]
		if isValidPath(xd, yd) && !vis[xd][yd] && word[i] == board[xd][yd] {
			vis[xd][yd] = true
			if helper(board, word, vis, xd, yd, i) {
				return true
			}
			vis[xd][yd] = false
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	rows = len(board)
	cols = len(board[0])
	vis := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		vis[i] = make([]bool, cols)
	}

	for i, col := range board {
		for j, v := range col {
			if v == word[0] {
				vis[i][j] = true
				if helper(board, word, vis, 0, i, j) {
					return true
				}
				vis[i][j] = false
			}
		}
	}
	return false
}
