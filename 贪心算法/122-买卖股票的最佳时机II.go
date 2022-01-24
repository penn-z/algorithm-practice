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

func main() {

}

// 贪心策略
func maxProfit(prices []int) int {
	/*
		题意可以无限买卖股票
		1. 只要发现当天股票比上一天大，则卖出；比下一天股票小，则买入
		2. 假设tmpProfit利润为第i-1天买入和第i天卖出后赚取的利润，则tmpProfit = prices[i] - prices[i-1]
		3. 若当天利润 > 0，则计入总收益，否则不计入（即不进行股票操作）
	*/

	if len(prices) <= 1 {
		return 0
	}

	sumProfit := 0
	for i := 1; i < len(prices); i++ {
		tmpProfit := prices[i] - prices[i-1]
		if tmpProfit > 0 {
			sumProfit += tmpProfit
		}
	}

	return sumProfit
}
