/*
	leetcode 300:
	给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
	子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/
package main

func main() {

}

// 动态规划
func lengthOfLISV2(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}

	// 1. 确定dp[i]数组及下标含义

	// 2. 确定递推公式
	/*
		dp[i]表示以nums[i]结尾的最大上升子序列
		则 dp[i] = max(dp[j]+1)，其中 0 <= j <= i-1， nums[j] < nums[i]
	*/

	// 3. dp初始化
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// 至少是其本身长度
		dp[i] = 1
	}

	// 4. 确定遍历顺序
	/*
		for i : 0 -> len(nums):
			for j : 0 -> i-1:
	*/
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i-1; j++ {
			// 可以组成以nums[i]结尾的最长上升子序列
			if nums[i] > nums[j] {
				// max{dp[0], dp[1], ... dp[i-1]} + 1, 这里遍历j时，每次与dp[i]比较，取较大值赋值后重新赋值给dp[i]
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	for i := range dp {
		if res < dp[i] {
			res = dp[i]
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
