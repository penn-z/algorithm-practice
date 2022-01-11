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
