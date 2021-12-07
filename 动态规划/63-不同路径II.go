/**
leetcode 63 - 不同路径II
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
	网格中的障碍物和空位置分别用 1 和 0 来表示。

思路: dp
1. 确定dp数组及其下标含义
	dp[i][j]: 到达(i, j)时，可能出现的不同路径
2. 确定递推公式
	在(i, j)时，可知此时不同路径数量为(i, j-1), (i-1, j)两者不同路径之和
	dp[i][j] = dp[i][j-1] + dp[i-1][j]

	if dp[i][j] 有障碍物 {
		dp[i][j] = 0
	} else {
		dp[i][j] = dp[i][j-1] + dp[i-1][j]
	}
3. dp数组初始化
	dp[0][0] = 0
	从(0, 0)走到(i, 0)只有1种走法, 故dp[i][0] = 1
	从(0, 0)走到(0, j)只有1种走法，故dp[0][j] = 1
4. 确定遍历顺序
	i, j分别顺序递增
5. 举例dp数组，与程序运行是否一致
*/
package main

import "fmt"

func main() {
	grid := [][]int{}
	for i := 0; i <= 2; i++ {
		arr := []int{}
		grid = append(grid, arr)
	}

	grid[0] = []int{0, 1}
	grid[1] = []int{1, 1}
	grid[2] = []int{0, 0}
	res := uniquePathsWithObstacles(grid)
	fmt.Println("===res: ", res)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 机器人从(0, 0)出发，到(m-1, n-1)终点
	if len(obstacleGrid) == 0 {
		return 0
	}

	lenM := len(obstacleGrid)
	lenN := len(obstacleGrid[0])

	// 递推公式: dp[i][j] = dp[i][j-1] + dp[i-1][j]
	dp := make([][]int, lenM)
	for i := range dp {
		dp[i] = make([]int, lenN)
	}

	// obstacleGrid[i][0] == 0有障碍，则dp[i][0] = 0
	for i := 0; i < lenM && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	// obstacleGrid[0][j] == 0有障碍，则dp[0][j] = 0
	for j := 0; j < lenN && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < lenM; i++ {
		for j := 1; j < lenN; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				// continue
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	// debug遍历dp
	// for i, _ := range dp {
	// 	for j, _ := range dp[i] {
	// 		fmt.Printf("%.2v,", dp[i][j])
	// 		// fmt.Printf("%+v,", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }

	return dp[lenM-1][lenN-1]
}
