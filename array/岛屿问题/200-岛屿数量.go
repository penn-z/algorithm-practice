/*
	leetcode 200: 岛屿数量
		给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
		岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
		此外，你可以假设该网格的四条边均被水包围。

	demo1:
		输入：grid = [
			["1","1","1","1","0"],
			["1","1","0","1","0"],
			["1","1","0","0","0"],
			["0","0","0","0","0"]
		]
		输出：1


	demo2:
		输入：grid = [
			["1","1","0","0","0"],
			["1","1","0","0","0"],
			["0","0","1","0","0"],
			["0","0","0","1","1"]
		]
		输出：3

	思路1: DFS递归网格
		1. 假设行坐标为r, 列坐标为c
		2. 递归base case，不能让网格越界
		3. 确定当前网格 上(r-1, c), 下(r+1, c), 左(r, c-1), 右(r, c+1) 网格是否岛屿
		4. 0表示非岛屿，1表示岛屿(未遍历)，2表示岛屿(已遍历)

*/

package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	// res := numIslands(grid)
	res := numIslandsV2(grid)
	fmt.Println("res: ", res)
}

// 思路1: DFS递归
func numIslands(grid [][]byte) int {
	res := 0

	m, n := len(grid), len(grid[0])
	// 遍历grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 每发现一个岛屿，则res++
				res++
				dfs(grid, i, j)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, r, c int) {
	m, n := len(grid), len(grid[0])

	// base case
	// 越变网格图边界
	if r < 0 || c < 0 || r >= m || c >= n {
		// 越界
		return
	}

	// 该网格不是岛屿，直接返回
	if grid[r][c] != '1' {
		return
	}

	grid[r][c] = '2' // 将网格标记为2，表示为岛屿且已访问

	// 访问上，下，左，右网格
	dfs(grid, r-1, c) // 上
	dfs(grid, r+1, c) // 下
	dfs(grid, r, c-1) // 左
	dfs(grid, r, c+1) // 右
}

// 思路2: BFS
func numIslandsV2(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				bfs(grid, i, j)
			}
		}
	}

	return res
}

// 深度优先遍历
func bfs(grid [][]byte, r, c int) {
	queue := [][]int{}
	queue = append(queue, []int{r, c})
	for len(queue) > 0 {
		// 队头元素出队
		head := queue[0]
		queue = queue[1:]
		i, j := head[0], head[1]
		// 越界
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			continue
		}

		if grid[i][j] != '1' {
			continue
		}

		grid[i][j] = '2' // 标记为岛屿(已访问)
		// 上下左右网格入队
		queue = append(queue,
			[]int{i - 1, j},
			[]int{i + 1, j},
			[]int{i, j - 1},
			[]int{i, j + 1},
		)
	}
}
