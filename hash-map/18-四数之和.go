/*
	leetcode 18. 四数之和
	给你一个由 n 个整数组成的数组 nums, 和一个目标值 target。请你找出并返回满足下述条件切不重复的四元组[nums[a], nums[b], nums[c], nums[d]] (若两个四元组元素一一对应，则认为两个四元组重复):
		* 0 <= a, b, c, d < n
		* a, b, c和d互不相同
		* nums[a] + nums[b] + nums[c] + nums[c] == target
	可以按任意顺序返回答案。

	demo 1:
		input: nums = [1, 0, -1, 0, -2, 2], target = 0
		output: [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]

	demo 2:
		input: nums = [2, 2, 2, 2, 2], target = 8
		output: [2, 2, 2, 2]

	思路: 双指针法，解法与15-三数之和类似
	1. 先把数组转换为升序数组
	2. 对数组进行两层遍历，外层i作为一指针，内层j亦作为一指针，left = j + 1, right = len(nums) - 1
	3. 注意两层遍历时的区中
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	res := fourSum(nums, 0)
	fmt.Println("res: ", res)
}

func fourSum(nums []int, target int) [][]int {
	// base
	if len(nums) < 4 {
		return [][]int{}
	}

	// 数组升序排序
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// 去重
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			// 迭代移动双指针
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					// 去重
					for left < right && nums[left] == nums[left+1] {
						left++
					}

					// 去重
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					// 移动双指针
					left++
					right--
				} else if sum < target {
					// sum比目标值小，left指针向右移
					left++
				} else if sum > target {
					// sum比目标值大，right指针向左移
					right--
				}
			}
		}
	}

	return res
}
