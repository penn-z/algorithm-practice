/*
	leetcode 153: 寻找旋转排序数组中的最小值

    思路1: 遍历
        1. 因为原本数组是升序排列，故而可以先假设nums[0]为最小值，遍历比较个元素大小
        2. 若出现小于min值的元素，可知该元素为最小值，结束遍历

	思路2: 二分查找
		1. 创建left, right指针分别指向nums的首尾数字，然后计算出middle，分类讨论:
			1. mid > right, 说明最小值在mid右侧，故left移动到mid+1位置
			2. mid < left, 说明最小值在mid左侧或者就是mid，故right移动到mid处
			3. mid == right，说明区间仅剩一个长度，即是最小值，直接返回nums[right]即可
*/

package main

func main() {

}

// 思路1: 暴力搜索
func findMin(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
			break
		}
	}

	return min
}

// 思路2: 二分查找
func findMinV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	for left < right { // 要保证区间长度大于1
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			// 最小值在mid右侧，故left移动到mid位置
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// 最小值在左侧
			right = mid
		} else {
			right--
		}
	}

	return nums[left]
}
