/*
	leetcode 695: 岛屿的最大面积

	给你一个大小为 m x n 的二进制矩阵 grid 。
	岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
	岛屿的面积是岛上值为 1 的单元格的数目。
	计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。


	思路1: dfs递归求解，在leetcode 200的基础上，递归时，每次面积+1即可
*/
package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := areaDFS(grid, i, j)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}

// 递归求解面积
func areaDFS(grid [][]int, r, c int) int {
	m, n := len(grid), len(grid[0])
	// base case
	if r < 0 || c < 0 || r >= m || c >= n {
		return 0
	}

	// 非未访问过的岛屿
	if grid[r][c] != 1 {
		return 0
	}

	// 当前网格标记为已访问过的岛屿
	grid[r][c] = 2

	// 递归求解当前网格上下左右网格面积之和

	return 1 +
		areaDFS(grid, r-1, c) + // 上
		areaDFS(grid, r+1, c) + // 下
		areaDFS(grid, r, c-1) + // 左
		areaDFS(grid, r, c+1) // 右
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
