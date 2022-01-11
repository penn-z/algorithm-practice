package main

func main() {

}

/*
	思路1: 直接遍历二维网格
		遇到是陆地时，判断当前格子的四边：
			若是与水域相邻，则该边计算周长，+1
			若是网格最外边界，计算该边周长，+1
			若是与其他陆地相邻，则不计
*/
func islandPerimeter(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				perimeter := calcPerimeter(grid, i, j)
				res += perimeter
			}
		}
	}

	return res
}

func calcPerimeter(grid [][]int, r, c int) int {
	perimeter := 0
	// 判断上边界
	if r-1 < 0 || grid[r-1][c] == 0 {
		perimeter++
	}

	// 下边界
	if r+1 >= len(grid) || grid[r+1][c] == 0 {
		perimeter++
	}

	// 左边界
	if c-1 < 0 || grid[r][c-1] == 0 {
		perimeter++
	}

	// 右边界
	if c+1 >= len(grid[0]) || grid[r][c+1] == 0 {
		perimeter++
	}

	return perimeter
}

/*
	思路2: dfs递归

	大体思路同leetcode-200， 不同之处在于处理当前格子 判断相邻上下左右网格时，
		如果越界，当前格子周长需要+1；
		如果当前格子是水域，则周长需要+1;
		如果是已遍历过的陆地，则不计算，跳过；

*/
func islandPerimeterDFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 题意限制只有一个岛屿，则计算1次即可
			if grid[i][j] == 1 {
				return dfs(grid, i, j)
			}
		}
	}

	return 0
}

func dfs(grid [][]int, r, c int) int {
	m, n := len(grid), len(grid[0])
	// 当前格子越界，周长需要+1，为父函数的一条边长
	if r < 0 || c < 0 || r >= m || c >= n {
		return 1
	}

	// 当前格子为水域, 周长+1，为父函数的一条边长
	if grid[r][c] == 0 {
		return 1
	}

	// 当前格子已遍历，不计算
	if grid[r][c] != 1 {
		return 0
	}

	grid[r][c] = 2

	return dfs(grid, r-1, c) + // 上
		dfs(grid, r+1, c) + // 下
		dfs(grid, r, c-1) + // 左
		dfs(grid, r, c+1) // 右
}
