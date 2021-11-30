/*
	leetcode 46: 全排列
	输入一组不重复的数字，返回它们的全排列

	demo1:
		输入：nums = [1,2,3]
		输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

	demo2:
		输入: nums = [0, 1]
		输出: [[0, 1], [1, 0]]
*/
package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3}
	nums := []int{1, 2, 3, 4}
	// nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Println("res: ", res)

}

var res = [][]int{}

func permute(nums []int) [][]int {
	res = [][]int{}
	if len(nums) == 0 {
		return res
	}

	// 记录结果路径
	path := []int{}
	// 记录当前路径是否适用过， map[index]bool
	used := make(map[int]bool)
	backtrack(nums, path, used)
	return res
}

func backtrack(nums, path []int, used map[int]bool) {
	// 结束条件
	if len(nums) == len(path) {
		// 加入结果中
		p := make([]int, len(path))
		// slice底层是指针引用，所以需要copy一份
		copy(p, path)
		res = append(res, p)
		return
	}

	for i, num := range nums {
		// 排除不合法的选择
		if used[i] {
			continue
		}

		// 做选择

		// 将该选择从选择列表移除
		used[i] = true
		// 路径.add(选择)
		path = append(path, num)

		// 进入下一层决策树
		backtrack(nums, path, used)

		// 回撤取消选择
		path = path[:len(path)-1]
		// 将该选择再加入选择列表
		used[i] = false
	}

	return
}
