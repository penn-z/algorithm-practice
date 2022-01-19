/*
	leetcode-494: 目标和
	给你一个整数数组 nums 和一个整数 target 。
	向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
	例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
	返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
*/
package main

import "fmt"

func main() {
	//
	nums := []int{1, 1, 1, 1, 1}
	// res := findTargetSumWays(nums, 3)
	res := findTargetSumWaysV2(nums, 3)
	fmt.Println("res: ", res)
}

func findTargetSumWays(nums []int, target int) int {
	/*
		仔细分析，其实是0-1背包问题:
			1. 其实就是把数组分为两堆，一堆plusSet为加法集合，一堆subSet为减法集合， sum = plusSet + subSet
			2. 可得到表达式为 sum - plusSet - plusSet = target，plusSet = (sum - target) / 2
				因为nums数组中元素都为非负整数，则plusSet也为非负整数，则等式成立条件为 sum - target为非负偶数
			3. 上述成立的话，问题转换为在数组nums中取若干元素，使得这些元素之和等于plusSet，计算选取元素的方案。

			4. 确定动态规划数组含义，dp[i][j]表示数组nums在前i个数中选取元素，使得这些元素之和等于j的方案数。
			5. dp[i][j]递推公式:
				dp[i][j] = {
					dp[i-1][j], j < nums[i],  // 如果 j < nums[i], 则不能选取nums[i], 此时有dp[i][j] = dp[i-1][j]
					dp[i-1][j] + dp[i-1][j - nums[i]], j >= nums[i]  // 如果j >= nums[i], 如果不选nums[i], 方案数是dp[i-1][j-1]; 如果选nums[i]，方案数是dp[i-1][j - nums[i]]
				}
			6. 初始化dp
				当没有任何元素可以选取时，
				dp[0][j] = {
					1, j = 0
					0, j >= 1
				}
	*/

	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	diff := sum - target
	// diff要为非负偶数
	if diff < 0 || diff%2 != 0 {
		return 0
	}

	plus := diff / 2

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, plus+1)
	}

	// 初始化dp
	dp[0][0] = 1

	// 当 1 <= i <= n时，第i个元素(i从0开始计数，即为nums[i-1])满足
	/*
		dp[i][j] = {
			dp[i-1][j], // j < nums[i-1]
			dp[i-1][j] + dp[i-1][j-nums[i-1]] // j >= nums[i-1]
		}
	*/
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= plus; j++ {
			if j < nums[i-1] {
				// 无法装入nums[i-1]
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入nums[i-1]方案+不装入nums[i-1]方案
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}

	fmt.Println("dp: ", dp)

	return dp[len(nums)][plus]
}

func findTargetSumWaysV2(nums []int, target int) int {
	/*
		公式推导：
		把nums分为plus，sub两个集合之和，plus为加法集合，sub为减法集合，则有
		sum = plus + sub
		target = sum - plus - plus = sum - 2*plus
		plus = (sum-target)/2
		可知，因为plus是非负整数，所以sum-target必须为正偶数整数

		1. 确认dp数组及其下标含义，dp[j]表示填满背包容量j时，有dp[j]种方法
		2. dp递推公式:
			dp[j] = dp[j] + dp[j - nums[i]]
			表示当前填满容量j的dp[j]方法 = 之前填满容量j的方法 + 之前填满j-nums[i]容量的方法
		3. dp初始化
			dp[0] = 1，表示当背包容量为0时，有1种填充方法，那就是不填
		4. 遍历顺序
			先遍历物品，后倒序遍历背包
		5. 结果即为dp[plus]
	*/
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	diff := sum - target
	// diff为非负偶整数
	if diff < 0 || diff%2 != 0 {
		return 0
	}

	plus := diff / 2

	dp := make([]int, plus+1)
	// 背包容量为0时，不填充即为填充方法
	dp[0] = 1
	// 遍历物品
	for i := 0; i < len(nums); i++ {
		// 倒序遍历背包
		for j := plus; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}

	return dp[plus]
}
