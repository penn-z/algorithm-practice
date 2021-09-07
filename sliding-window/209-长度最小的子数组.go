/*
	给定一个含有 n 个正整数的数组和一个正整数 target 。
	找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0

	demo1:
		输入：target = 7, nums = [2,3,1,2,4,3]
		输出：2
		解释：子数组 [4,3] 是该条件下的长度最小的子数组。


	demo2:
		输入：target = 4, nums = [1,4,4]
		输出：1

	思路:
		1. 使用滑动窗口思想，滑动窗口左右边界，不断维护窗口内的数据

*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	res := minSubArrayLen(target, nums)
	fmt.Printf("====res:%d\n", res)
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 窗口内的值之和
	windowsSum := nums[0]

	// base case
	if windowsSum >= target {
		return 1
	}

	// 符合题意的最小长度
	minLen := 0

	fmt.Printf("array: %+v\n", nums)

	// 窗口左右边界值
	left, right := 0, 0
	for right < len(nums)-1 {
		// 右移窗口右边界
		right++

		fmt.Printf("add right ele:%d\n", nums[right])

		// 进行窗口数据更新
		windowsSum += nums[right]
		if windowsSum >= target {
			minLen = getMinValue(minLen, right-left+1)
			fmt.Printf(">>>>minLen:%d\n", minLen)
		}

		/**debug**/
		fmt.Printf("windows left:%d, right:%d, arr:%+v\n", left, right, nums[left:right+1])

		// 判断左窗口是否需要收缩
		for windowsSum > target {
			// 移除元素
			deleteNums := nums[left]
			fmt.Printf("delete left index:%d, ele:%d\n", left, deleteNums)

			// 左移窗口左边界，收缩窗口
			left++

			// 更新窗口数据
			windowsSum -= deleteNums
			if windowsSum >= target {
				minLen = getMinValue(minLen, right-left+1)
				fmt.Printf("<<<<minLen:%d\n", minLen)
			}
		}
	}

	return minLen
}

func getMinValue(i, j int) int {
	if i == 0 {
		return j
	}

	if i < j {
		return i
	}

	return j
}
