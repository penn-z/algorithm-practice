/*
	leetcode 674 -最长连续递增序列
	给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
	连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

*/
package main

func main() {

}

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// 1. 确认dp数组及其下标含义
	// dp[i]表示数组nums第i-1元素的最长递增子序列长度

	// 2. 递推公式
	// dp[i] = dp[i-1] + 1 且 nums[i] > nums[i-1]

	// 3. dp初始化
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	// 4. 确定顺序
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}

	res := 0
	for i := range dp {
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}
