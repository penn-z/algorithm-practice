/**
leetcode 188 - 买卖股票的最佳时机IV
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



思路: 动态规划
1. dp数组及其含义
	需要定义三个状态：天数i; 是否持有股票j: 0未持有，1持有; 已交易次数k，卖出时才算是完成一次交易
	则 dp[i][j][k]表示第i天，是否持有股票，卖出次数为k的收益
2. 确定递推公式
	dp[i][0][k] = max{ dp[i-1][0][k](无交易) + dp[i-1][1][k-1]+prices[i] （卖出）}
	dp[i][1][k] = max{ dp[i-1][1][k](无交易) + dp[i-1][0][k]-prices[i](买入) }
3. 确定dp初始化值
	[天数][是否持有股票][已卖出次数]
	第一天持有股票:
		dp[0][1][0] = -prices[0]
		其实无论第一天卖出多少次，只要持有股票，收益都是 -prices[0]，也即
		dp[0][1][k] = -prices[0]

	第一天不持有股票:
		dp[0][0][0] = 0
		无论第一天卖出多少次，只要不持有股票，收益都是0，即
		dp[0][0][k] = 0
4. 确定遍历顺序
	外层
	for i := 1; i < len(prices) ; i++ {
		for j := 0; j <= k; j++ {
			// 未持有股票
			dp[i][0][j] = max(dp[i-1][0][j], dp[i-1][1][j-1] + prices[i]) // 未交易, 卖出

			// 持有股票
			dp[i][1][j] = max(dp[i-1][1][j], dp[i-1][0][j] - prices[i]) // 未交易, 买入
		}
	}



*/
package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfitV2(2, prices)
	fmt.Println("===res: ", res)
}

func maxProfitV2(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// dp状态
	// 未持有股票
	// dp[i][0][j] = max(dp[i-1][0][j], dp[i-1][1][j-1] + prices[i])

	// 持有股票
	// dp[i][1][j] = max(dp[i-1][1][j], dp[i-1][0][j] - prices[i])

	// 创建三维数组
	// 三个状态: [天数][是否持有][卖出次数]
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, k+1)
		}
	}

	// dp数组初始化
	for i := 0; i <= k; i++ {
		// 第一天就买入
		dp[0][1][i] = -prices[0]
		// 第一天未买入
		dp[0][0][i] = 0
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < k+1; j++ {
			if j != 0 { // 防止数组越界
				// 未持有股票
				dp[i][0][j] = max(dp[i-1][0][j], dp[i-1][1][j-1]+prices[i])
			}

			// 持有股票
			dp[i][1][j] = max(dp[i-1][1][j], dp[i-1][0][j]-prices[i])
		}
	}

	return dp[len(prices)-1][0][k]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
