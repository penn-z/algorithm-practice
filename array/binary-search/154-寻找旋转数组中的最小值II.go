/*
	leetcode 154: 寻找旋转数组中的最小值II
	已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
	若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
	若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
	注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

	给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

	思路1: 暴力搜索
		1. 遍历数组，暴力搜索即可


	思路2: 二分查找
		1. 创建left, right指针分别指向nums的首尾数字，然后计算出middle，分类讨论:
			1. mid > right, 说明最小值在mid右侧，故left移动到mid+1位置
			2. mid < left, 说明最小值在mid左侧或者就是mid，故right移动到mid处
			3. mid == right时，因为有重复元素存在，我们没法知道min是在mid左侧还是右侧，那无论right是不是最小值，都有替代品mid可以替代，故而right--再次循环即可
*/
package main

func main() {

}

func findMin(nums []int) int {
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
