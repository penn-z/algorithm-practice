/*
	leetcode 34: 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
	如果数组中不存在目标值 target，返回 [-1, -1]。

	思路:
		1. 可知题意是求解二分查找左侧边界与右侧边界
		2. 左右边界各来一次即可
*/
package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	res := searchRange(nums, target)
	fmt.Printf("=====res:%+v\n", res)
}

func searchRange(nums []int, target int) []int {
	// base case
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftBound := left_bound(nums, target)
	rightBound := right_bound(nums, target)
	// rightBound := -1

	return []int{leftBound, rightBound}
}

// 搜索左侧边界
func left_bound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 使用左闭右开
	left := 0
	right := len(nums)
	// 左闭右开，当left == right时退出，故而 使用 <
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			//不返回mid，收缩右侧边界，搜索左边区间
			right = mid
		} else if nums[mid] < target {
			// 目标结果在右边区间，收缩左侧边界
			left = mid + 1
		} else if nums[mid] > target {
			// 目标结果在左边区间，收缩右侧边界
			right = mid
		}
	}

	// 越界，target比所有数大
	if left == len(nums) {
		return -1
	}

	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

// 搜索右侧边界
func right_bound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 使用左闭右开区间[left, right)
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			// 不急于返回mid，收缩左侧边界
			left = mid + 1
		} else if nums[mid] < target {
			// 目标结果在右边区间，收缩左侧边界
			left = mid + 1
		} else if nums[mid] > target {
			// 目标结果在左边区间，收缩右侧边界
			right = mid
		}
	}

	// 终止条件left == right，若存在目标，left > 0
	if left == 0 {
		return -1
	}

	// 因为left = mid + 1，此时nums[left]必不等于target, nums[left-1]有可能相等
	if nums[left-1] == target {
		return left - 1
	} else {
		return -1
	}
}
