package main

import "fmt"

func main() {
	n := 5
	res := climbStairs(n)
	fmt.Println("res:", res)
}

/**
思路: 动态规划
1. 确定dp数组及其下标的含义
	dp[i]: 爬到第i层楼梯，有dp[i]种方法

2. 确定递推公式
	dp[i] = dp[i-1] + dp[i-2]

3. dp数组如何初始化
	dp[0] = 1
	dp[1] = 1, dp[2] = 2, 从i = 3开始递推
4. 确定遍历顺序
	从dp[i] = dp[i-1] + dp[i-2] 可知，从前往后遍历
5. 举例dp数组
	i = 1, dp[1] = 1
	i = 2, dp[2] = 2
	i = 3, dp[3] = 3
	i = 4, dp[4] = 5
	i = 5, dp[5] = 8
*/
func climbStairs(n int) int {
	// dp[i] = dp[i-1] + dp[i-2]

	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
