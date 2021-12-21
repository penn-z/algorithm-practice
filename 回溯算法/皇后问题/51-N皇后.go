/*
	leetcode 51: N皇后问题
	i为纵坐标，j为横坐标
	确定一个queen位置(i, j)后，则其所在
		行 (x1, j)
		列 (i, x2)
		左斜线 (x3, y3) => x3 + y3 == i + j
		右斜线 (x4, y4) => x4 - y4 == i - j
	均不可以出现其他queens，可以用三个set分别表示列，左斜线，右斜线是否出现过queens

	1. 初始化棋盘board，每个位置都填充为"."
	2. 初始化列，左斜线，右斜线三个set
	3. 开始进行深度递归回溯
	4. 回溯函数中:
		1. 若已经来到最后一行，则达到结束条件，需要记录下当前棋盘的queens放置位置
		2. 未到结束条件，则遍历列
			1. 判断同列，同左右斜线是否queen
			2. 无queen则把列、左右斜线标记上有queen存在，该位置标记为"Q"，递归调用回溯函数
			3. 回溯完成，则撤销刚才的标记操作
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 4
	res := solveNQueens(n)
	fmt.Println("last res: ", res)
}

var (
	res        = [][]string{}
	cols       = map[int]bool{}
	leftSlash  = map[int]bool{}
	rightSlash = map[int]bool{}
)

func solveNQueens(n int) [][]string {
	res = [][]string{}
	if n < 1 {
		return res
	}

	// 初始化棋盘，全部置为.
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	// cols, leftSlash, rightSlash => set
	cols, leftSlash, rightSlash = make(map[int]bool), make(map[int]bool), make(map[int]bool)

	backtrack(n, 0, board)

	return res
}

func backtrack(n, row int, board [][]string) {
	// 停止条件，已经扫完最后一行
	if n == row {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			// 每一行的棋盘都取出
			temp[i] = strings.Join(board[i], "")
		}

		res = append(res, temp)

		return
	}

	// 选择 in 选择列表
	for c := 0; c < n; c++ {
		// i 为列索引

		// 排除不合法选项
		// 同一列已有queen
		if cols[c] {
			continue
		}

		// 同一左斜线已有queen
		leftSlashIndex := row + c
		if leftSlash[leftSlashIndex] {
			continue
		}

		// 同一右斜线已有queen
		rightSlashIndex := row - c
		if rightSlash[rightSlashIndex] {
			continue
		}

		// 做选择
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
