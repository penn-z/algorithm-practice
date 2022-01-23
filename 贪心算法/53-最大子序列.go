/*
	leetcode 53: 最大子序列之和
	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
	子数组 是数组中的一个连续部分。

	demo1:
		输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
		输出：6
		解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

	思路1: 动态规划
		1. 确定dp[i]数组及含义:
			dp[i]表示包含第i个元素的连续子数组最大和
		2. 确定dp[i]递推公式
			dp[i] = max {
				dp[i-1] + nums[i], // dp[i-1] >= 0, 前面的连续子数组之和为整数，则nums[i]与其连成一片
				nums[i], // dp[i-1] < 0，前面的连续子数组之和为负数，nums[i]自成一派
			}
		3. dp初始化
		4. 确定dp遍历顺序
*/
package main

func main() {

}

/*
思路1: 动态规划
	1. 确定dp[i]数组及含义:
		dp[i]表示包含第i个元素的连续子数组最大和
	2. 确定dp[i]递推公式
		dp[i] = max {
			dp[i-1] + nums[i], // dp[i-1] >= 0, 前面的连续子数组之和为整数，则nums[i]与其连成一片
			nums[i], // dp[i-1] < 0，前面的连续子数组之和为负数，nums[i]自成一派
		}
	3. dp初始化
	4. 确定dp遍历顺序
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	maxInt := int(^uint(0) >> 1)
	minInt := int(^maxInt)
	res := minInt
	for i := range dp {
		res = max(res, dp[i])
	}

	return res
}

/*
	贪心策略：
		1. 用sum记录当前子序列之和，迭代过程中如果小于0，则放弃
		2. res记录最大的子序列之和
*/
func maxSubArrayV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxInt := int(^uint(0) >> 1)
	minInt := int(^maxInt)
	sum, res := 0, minInt
	for i := range nums {
		sum += nums[i]
		res = max(res, sum)
		if sum < 0 {
			sum = 0
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
