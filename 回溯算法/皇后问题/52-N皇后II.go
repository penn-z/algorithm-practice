package main

func main() {

}

var (
	res        int
	cols       = map[int]bool{}
	leftSlash  = map[int]bool{}
	rightSlash = map[int]bool{}
)

func totalNQueens(n int) int {
	res = 0
	if n < 1 {
		return res
	}

	// 初始化棋盘, 全部置为"."
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	// 初始化列、左斜线、右斜线 set
	cols, leftSlash, rightSlash = make(map[int]bool), make(map[int]bool), make(map[int]bool)

	backtrack(n, 0, board)
	return res
}

func backtrack(n, row int, board [][]string) {
	// 结束条件
	if n == row {
		// 已经到最后一行
		res++
	}

	// 选择 in 选择列表
	for c := 0; c < n; c++ {
		// 排除不合法选项
		// 左斜线
		leftSlashIndex := row + c
		// 右斜线
		rightSlashIndex := row - c

		if cols[c] || leftSlash[leftSlashIndex] || rightSlash[rightSlashIndex] {
			continue
		}

		// 选择当前位置为queen
		board[row][c] = "Q"
		cols[c] = true
		leftSlash[leftSlashIndex] = true
		rightSlash[rightSlashIndex] = true

		// 回溯
		backtrack(n, row+1, board)

		// 撤销选择
		board[row][c] = "."
		cols[c] = false
		leftSlash[leftSlashIndex] = false
		rightSlash[rightSlashIndex] = false
	}
}
