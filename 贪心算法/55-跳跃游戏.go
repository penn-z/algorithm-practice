/*
	leetcode 55: 跳跃游戏
	给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个下标。

	demo1:
		输入：nums = [2,3,1,1,4]
		输出：true
		解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

	demo2:
		输入：nums = [3,2,1,0,4]
		输出：false
		解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
*/

package main

func main() {

}

func canJump(nums []int) bool {
	/*
	   贪心策略:
	       1. 当前元素值是一个覆盖范围，每次移动更新覆盖范围，如果覆盖范围小于0，则结束循环
	*/

	coverRange := 0
	terminalLen := len(nums) - 1
	for i, v := range nums {
		if coverRange < v {
			// 更新覆盖范围
			coverRange = v
		}

		diffFromTerinal := terminalLen - i - 1
		if diffFromTerinal < coverRange {
			// 可达终点
			return true
		}

		// 覆盖范围已不能达到终点
		if coverRange < 1 {
			return false
		}

		// 当前不可达，遍历前进
		coverRange--
	}

	return false
}
