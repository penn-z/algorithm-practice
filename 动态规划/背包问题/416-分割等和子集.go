/*
	给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

	demo1:
		输入: nums = [1, 5, 11, 5]
		输出: true
		解释：数组可以分割成 [1, 5, 5] 和 [11] 。

	demo2:
		输入：nums = [1,2,3,5]
		输出：false
		解释：数组不能分割成两个元素和相等的子集。



*/

package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 7, 4}
	res := canPartition(nums)
	fmt.Println("res:", res)
	resV2 := canPartitionNotCompress(nums)
	fmt.Println("resV2:", resV2)
}

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	/*
		细看之下，就是0-1背包问题，背包容量为sum/2，物品是nums[i]，重量是nums[i]，物品价值也是nums[i]。
		则压缩后的dp为dp[j]，表示背包容量为j时，最大可以凑成子集总和为dp[j]

		1. dp[j]含义： 表示背包容量为j时，最大可以凑成子集总和为dp[j]
		2. 递推dp公式：
			dp[j] = max{dp[j], dp[j - nums[i]] + nums[i]}
		3. dp初始化：
			dp[0] = 0，其他下标，如果nums中元素为非负整数，则dp[j] = 0，否则为非无穷大。
		4. 确定遍历
			// 先遍历物品
				// 后遍历容量
		5. 推导dp验证
		6. 返回最终结果
	*/

	// dp := make()
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		// 无法整除，GG
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	// 初始化，因为nums元素为非负整数，故所有dp[j]都为0

	// 遍历物品
	for i := 0; i < len(nums); i++ {
		// 遍历背包
		// 每个元素不可重复放入，故而需要从后往前遍历，以免重复覆盖
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	fmt.Println("v1 dp: ", dp)

	// 集合中的元素正好可以凑成总和target
	if dp[target] == target {
		return true
	}

	return false
}

// 0-1背包，二维状态不压缩
func canPartitionNotCompress(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	/*
		给一个容量为sum/2的背包，n个物品，每个物品的重量为nums[i]。现在装物品，是否存在一种装法，能够恰好装满背包？

		1. dp[i][j] = boolean表示，对于前i个物品，当前背包容量为j时，若dp[i][j] = true，则说明背包恰好可以装满，若dp[i][j] = false，说明不能恰好装满背包。

		2. 递推公式
			dp[i][j] = max {
				dp[i-1][j],	// 无法装入物品i
				dp[i-1][j - nums[i]], // 可以装入物品i，恰好装满背包
			}

		3. dp初始化
			dp[...][0] = true，表示背包没有空间时，已经装满
			dp[0][...] = false，当没有物品可以选择时，没办法装满背包

		4. 遍历顺序
			// 先遍历物品
				// 后遍历背包
	*/

	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}

	// 初始化
	for i := 0; i < len(nums); i++ {
		// 背包没有空间，已经恰好装满
		dp[i][0] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			// 背包容量不足，不能装入第i个物品
			if j-nums[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}

	fmt.Println("v2 dp: ", dp)

	return dp[len(nums)-1][target]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
