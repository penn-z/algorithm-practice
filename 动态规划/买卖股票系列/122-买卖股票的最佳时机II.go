/**
leetcode 122 - 买卖股票的最佳时机II
给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

demo1:
	input: [7, 1, 5, 3, 6, 4]
	output: 7
	explain: 第2天买入，第3天卖出, 该笔利润profit1 = 5-1 = 4
		第4天买入，第5天卖出，该笔利润profit2 = 6-3 = 3
		则总利润为 profit = profit1 + profit2 = 4 + 3 = 7
*/
package main

import "fmt"

func main() {
	test := int(^uint(0) >> 1)
	min := ^test
	fmt.Printf("test:%d, min:%d\n", test, min)
}

/**
思路: 动态规划
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	/**
	1. 确定dp数组及其下标含义
	dp[i][j]为第i天，是否持有股票，j取值范围[0, 1]
	*/
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	/**
	2. 确定递推公式
	未持有股票时:
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])  (不动，卖出)
	持有股票时:
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])  (不动，买入)
	*/

	// 3. 初始化dp数组
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// 4. 确定遍历顺序
	for i := 1; i < len(prices); i++ {
		// 未持有股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // (不动，卖出)
		// 持有股票
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // (不动，买入)
	}

	// 最终结果
	return dp[len(prices)-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

/**
思路: 贪心算法，遍历数组
1. 因为可以无限次买入卖出，则只要发现当前元素比上一个元素大，则卖出；比下一个元素小，则买入，以此达到收益最大化
2. 设tmp 为第i-1日买入与第i日卖出赚取的利润，即tmp = prices[i] - prices[i-1]
3, 若当天利润tmp > 0，则计入总收益，否则不计入总利润(及不进行股票操作)
*/
func maxProfitV2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// lowPrice, sumProfit := int(^uint(0)>>1), 0
	sumProfit := 0
	for i := 0; i < len(prices); i++ {
		var tmp int
		if i > 0 {
			tmp = prices[i] - prices[i-1]
		}

		if tmp > 0 {
			sumProfit += tmp
		}
	}

	return sumProfit
}
