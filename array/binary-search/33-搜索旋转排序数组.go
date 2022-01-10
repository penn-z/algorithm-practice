/*
	leetcode 33 - 搜索旋转排序数组

	整数数组 nums 按升序排列，数组中的值 互不相同 。
	在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
	给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1。

	demo1:
		输入：nums = [4,5,6,7,0,1,2], target = 0
		输出：4

	demo2:
		输入：nums = [4,5,6,7,0,1,2], target = 3
		输出：-1

*/
package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	res := search(nums, 5)
	fmt.Println(res)
}

// 暴力搜索, O(n)
func search1(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

/*
	二分查找
	由题意可知，旋转后的数组被分为左右两部分，总有一部分是单调递增的
	1. 可以通过nums[0]与nums[mid]比较大小，确定[l, mid]或者[mid+1, r]有序
	2. 然后查看target是否在有序范湖区间内，是则搜索这个区间，否则搜索另外半边区间
*/
func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		// 取中点
		mid := left + (right-left)/2
		// 中点就是目标值，直接返回
		if nums[mid] == target {
			return mid
		}

		// 基准点nums[0]确定mid哪半部分有单调增区间
		if nums[0] <= nums[mid] {
			// [0, mid)区间为单调增区间

			// 判断target是否在该增区间内
			if nums[0] <= target && target < nums[mid] {
				// 在有序区间[0, mid]中
				right = mid - 1
			} else {
				// 在另外一半区间中搜索
				left = mid + 1
			}
		} else {
			// [mid+1, right] 区间为单调区间

			// 判断target是否该增区间中
			if nums[mid] < target && target <= nums[len(nums)-1] {
				// 在有序区间[mid, right]中
				left = mid + 1
			} else {
				// 在另外一半区间中搜索
				right = mid - 1
			}
		}
	}

	return -1
}
