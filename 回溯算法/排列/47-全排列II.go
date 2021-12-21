/*
	leetcode 47:
		给定一个可包含重复数字的序列nums, 按任意顺序返回所有不重复的全排列。

	demo1:
		input: [1, 1, 2]
		output: [
			[1, 1, 2],
			[1, 2, 1],
			[2, 1, 1],
		]
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 2, 1}
	nums := []int{3, 3, 0, 3}
	ret := permuteUnique(nums)
	fmt.Println("res: ", ret)
}

var (
	res = [][]int{}
)

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	if len(nums) == 0 {
		return res
	}

	// 升序数组，使得相同的元素相邻
	sort.Ints(nums)

	// 选择的路径
	path := []int{}
	// used := map[int]bool{}
	used := make([]bool, len(nums))
	backtrack(nums, path, used)
	return res
}

func backtrack(nums []int, path []int, used []bool) {
	// 结束条件
	if len(nums) == len(path) {
		// 需要把path加入到结果集中
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
		return
	}

	// for 选择 in 选择列表 {}
	for i, num := range nums {
		// 排除不合法的路径
		if used[i] {
			continue
		}

		// used[i-1] == true不同层，nums[i]为nums[i-1]子节点，即使nums[i] == nums[i-1]也合法。

		// used[i-1] == false是因为nums[i-1]在回退的过程中刚被撤销选择，目前处于同一层
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}

		// 做选择, 路径add
		path = append(path, num)
		used[i] = true
		// 进入下一层决策
		backtrack(nums, path, used)

		// 核销选择
		path = path[:len(path)-1]
		used[i] = false
	}

	return
}
