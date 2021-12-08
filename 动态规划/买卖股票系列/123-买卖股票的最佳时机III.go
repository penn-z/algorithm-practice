package main

func main() {

}

func maxProfit(prices []int) int {
	// 假设最多可以交易次数为k，题意k=2
	var k int = 2
	// 确定dp数组及下标的含义
	/*
	   dp[天数][是否持有股票][卖出次数]
	   未持有股票时:
	       dp[i][0][k] = max(dp[i-1][0][k], dp[i-i][1][k-1]+prices[i]) (不动，卖出)
	   持有股票时:
	       dp[i][1][k] = max(dp[i-1][1][k], dp[i-1][0][k]-prices[i]) (不动, 买入)
	*/
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, k+1)
		}
	}

	// dp数组初始化
	for i := 0; i <= k; i++ {
		// 第一天买入
		dp[0][1][i] = -prices[0]
		// 第一天未买入
		dp[0][0][i] = 0
	}

	// 确定遍历顺序
	for i := 1; i < len(prices); i++ {
		for j := 0; j <= k; j++ {
			if j != 0 { // 防止数组越界
				// 未持有股票
				dp[i][0][j] = max(dp[i-1][0][j], dp[i-1][1][j-1]+prices[i]) // (不动, 卖出)
			}

			// 持有股票
			dp[i][1][j] = max(dp[i-1][1][j], dp[i-1][0][j]-prices[i]) // (不动, 买入)
		}
	}

	// 最终结果
	return dp[len(prices)-1][0][k]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
