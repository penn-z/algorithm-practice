/*
	0-1背包问题：
	假设有重量为[1, 3, 4]的三个物品，价值为[15, 20, 30]，背包重量为4，求背包能装下的最大价值是多少。

	思路: 动态规划
		1. 确定dp数组及其下标含义，dp[i][j]表示从[0-i]的物品里任意取，放进容量为j的背包中，价值总和的最大值.
		2. 确定递推公式，dp[i][j]可由
			dp[i-1][j], 不取物品i，因为取了物品i后，重量总和会大于j；
			dp[i-1][j - weight[i]] + value[i]，取物品i，取了物品i后，总重量不会超过j
		3. dp初始化
			dp[i][0]，表示背包容量j=0，则什么也装不下，dp[i][0] = 0
			dp[0][j]表示，装第一个物品，
				若j < weight[0]，第一个物品装不下，dp[0][j] = 0;
				若j >= weight[0]，区间为[weight[0], begWeight]，则dp[0][j] = value[0]

		4. 确定遍历顺序
			先遍历物品，后遍历背包
			for i := 1; i < len(weight); i++ { // 遍历物品
				for j := 0; j <= weight; j++ {

				}
			}


*/
package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	res := package0_1(weight, value, 4)
	fmt.Println("res: ", res)
}

func package0_1(weight, value []int, bagWeight int) int {
	if len(weight) == 0 {
		return 0
	}

	if bagWeight == 0 {
		return 0
	}

	/*
		1. 确定dp数组及下标含义：
			dp[i][j]表示从[0, j]的物品任意取，放进容量为j的背包中，价值总和的最大值.
	*/

	/*
		2. 递推dp[i][j]公式，
			dp[i][j] = max{
				dp[i-1][j], // 表示无法放入物品i，会超过背包容量j，此时背包中的总价值
				dp[i-1][j - weight[i]] + value[i] // 表示可以放入物品i，此时背包红的总价值
			}
	*/
	/*
		3. dp初始化
			当背包容量为0时，dp[i][0] = 0

			当取第一个物品时，dp[0][j]
				若 j < weight[0]，第一个物品无法放入
					dp[0][j] = 0
				若 j >= weight[0], 取值区间为[weight[0], begWegiht],
					dp[0][j] = value[0]
	*/

	/*
		4. 确定遍历顺序
			先遍历物品，后遍历背包
			for i := 1; i < len(weight); i++ {
				for j := 0; j <= begWeight; j++ {

				}
			}
	*/

	dp := make([][]int, len(weight))

	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	// dp初始化，当背包容量为0时，背包价值为0
	for i := 0; i < len(weight); i++ {
		dp[i][0] = 0
	}

	for j := 0; j <= bagWeight; j++ {
		if j < weight[0] {
			// 背包容量小于第一个物品weight
			dp[0][j] = 0
		} else {
			// 背包容量不小于第一个物品weight
			dp[0][j] = value[0]
		}
	}

	// 确定遍历顺序
	// 先遍历物品，后遍历背包
	for i := 1; i < len(weight); i++ {
		// 遍历背包
		for j := weight[i]; j <= bagWeight; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}

	return dp[len(weight)-1][bagWeight]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
