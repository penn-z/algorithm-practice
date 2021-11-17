/*
	leetcode 15: 三数之和
	给你一个包含 n 个整数的数组nums, 判断nums中是否存在三个元素a, b, c, 使得a + b + c = 0 ? 请你找出所有和为0且不重复的三元组。
	注意: 答案中不可包含重复的三元组。

	demo1:
		输入: nums = [-1, 0, 1, 2, -1, -4]
		输出: [[-1,-1,2], [-1, 0, 1]]

	demo2:
		输入: nums = []
		输出: []

	思路:
		双指针法
		1. len(nums) < 3, 直接返回
		2. 首先对数组进行升序排序
		3. 固定一个数nums[i]，再使用左右指针指向nums[left] = nums[i+1], nums[right] = nums[len(nums)-1]，就算三个数的和sum是否为0，满足则添加进结果集
		4. 如果nums[i] == nums[i-1]，说明该数字重复，会导致重复结果，所以跳过
		5. 当 sum == 0 时，如果 nums[left] == nums[left+1], 则会导致重复结果，应该跳过, left++
		6. 当 sum == 0 是，如果 nums[right] == nums[right-1]，会导致重复结果，应该跳过, right--
		7. 返回最终结果集即可

*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	// nums := []int{1, -1, -1, 0}
	// nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	// sort.Sort(sort.IntSlice(nums))
	fmt.Println("res", threeSum(nums))

}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	res := [][]int{}
	// 对数组进行排序
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	// 三指针, nums[i], nums[left], nums[right]
	for i := 0; i < len(nums); i++ {
		// i指针指向元素大于0，不符合题意，跳过
		if nums[i] > 0 {
			continue
		}

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1          // left指向i下一位
		right := len(nums) - 1 // right指向当前数组最后一位
		// 迭代移动双指针
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {

					// 去重
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					// 去重
					right--
				}

				// 收缩left, right
				left++
				right--
			} else if sum < 0 {
				// left右移
				left++
			} else if sum > 0 {
				// right向左移
				right--
			}

		}
	}

	return res
}
