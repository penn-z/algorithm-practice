/**
leetcode 746: 使用花费最小怕楼梯
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。
每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。
请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

demo1:
	input: cost = [10, 15, 20]
	output: 15
	explain: 最低花费从cost[1]开始，然后走两步即可达到阶梯顶，一共花费15

demo2:
	input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
	output: 6
	explain: 最低花费是从cost[0]开始，逐个经过那些1, 跳过cost[3]，一共花费6


*/
package main

import "fmt"

func main() {
	// cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	cost := []int{10, 15, 20}
	res := minCostClimbingStairs(cost)
	fmt.Println("==res:", res)

}

/**
思路: 动态规划
1. 确定dp数组及其下标含义
	dp[i]定义: 到达定i个台阶顶部时所花费的最小值dp[i]
2. 确定递推公式
	两个途径:
		1. 一个是dp[i-1]
		2. 一个是dp[i-2]
	从两个途径中选择最小值，算上当前i花费，则dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
3. dp数组如何初始化
	dp[0] = cost[0]
	dp[1] = cost[1]
4. 确定遍历顺序
	由题意可知是从前往后遍历
5. 举例dp数组，与程序运行的是否一致
*/
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	// 递推公式: dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	dp := make([]int, len(cost), len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	// 这里遍历少了最后一步
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
		fmt.Printf("===i:%d , dp:%+v\n", i, dp)
	}

	// 最后一步可以不用花费，所以取倒数第一、第二步的最小值
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
