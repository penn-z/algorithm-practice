package main

func main() {

}

// 动态规划
func maxSubArrayV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	/**
	  1. 确定dp数组及其下标含义
	  dp[i] = max(dp[i-1] + nums[i], nums[i])
	*/
	dp := make([]int, len(nums))

	/**
	  2. 确定递推公式
	  dp[i]两种情况:
	      1. 前面元素之和为整数，能使自身增大，则与前面连成一片
	      2. 前面元素之和为负数，则自己开辟新的连续数组
	  dp[i] = max(dp[i-1] + nums[i], nums[i])
	*/

	// 3. dp数组初始化
	dp[0] = nums[0]

	// 4. 确定数组遍历顺序
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	// 5. debug, 推导dp
	// fmt.Println("---dp: ", dp)

	// 最终结果
	maxInt := int(^uint(0) >> 1)
	res := int(^maxInt)
	for i := range dp {
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

/**
贪心算法
1. 遍历数组，依次相加求sum，若发现sum < 0, 则重新寻找子串
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxInt := int(^uint(0) >> 1)
	result := int(^maxInt)
	sum := 0
	for i := range nums {
		sum += nums[i]
		result = max(result, sum)
		if sum < 0 {
			sum = 0
		}
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
