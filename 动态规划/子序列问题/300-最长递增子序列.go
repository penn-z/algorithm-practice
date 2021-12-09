/*
	leetcode 300:
	给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
	子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/
package main

func main() {

}

// 动态规划
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// 1. 确定dp数组及其下标含义
	/*
	   dp[i]表示数组nums在下标为i时的最长递增子序列(可以不连续)
	*/

	// 2. 确定递推公式
	/*
	   dp[i] = max{dp[0], dp[1], ... dp[i-1]} + 1
	*/
	dp := make([]int, len(nums))

	// 3. dp数组初始化
	for i := range dp {
		// 初始状态下，每个元素至少是长度为1的递增子序列
		dp[i] = 1
	}

	// 4. 确定遍历顺序
	/*
	   for i : 1 -> (lenNums-1):
	       for j : 0 -> i-1 :
	           TODO
	*/
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= i-1; j++ {
			// 严格递增
			if nums[i] > nums[j] {
				// max{dp[0], dp[1], ... dp[i-1]} + 1, 这里遍历j时，每次与dp[i]比较，取较大值赋值后重新赋值给dp[i]
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 最终结果
	res := 0
	for i := range dp {
		res = max(res, dp[i])
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
