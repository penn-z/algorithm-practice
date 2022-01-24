/*
	leetcode-120: 三角形最小路径和

	给定一个三角形 triangle ，找出自顶向下的最小路径和。
	每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

	demo1:
		输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。


	思路1: 递归+备忘录记忆化
	若当前位置为(i, j)，i为纵坐标，j为横坐标
	定义f(i, j)为(i, j)点到底边的最小路径和，则递归公式为
	f(i, j) = min{f(i+1, j), f(i+1, j+1)} + triangle[i][j]
	由此，将任一点到底边的最小路径和，转化为了与该点相邻两点到底边的最小路径和中的较小值，再加上该点本身的值。


*/
package main

import "fmt"

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	res := minimumTotal(triangle)
	fmt.Println("res:", res)

}

func minimumTotal(triangle [][]int) int {
	depth := len(triangle)
	width := len(triangle[len(triangle)-1])
	memo := make([][]int, depth)
	for i := range memo {
		memo[i] = make([]int, width)
		for j := range memo[i] {
			memo[i][j] = -1000000
		}
	}

	return dfs(triangle, 0, 0, memo)
}

func dfs(triangle [][]int, i, j int, memo [][]int) int {
	if len(triangle) == i {
		// 越界
		return 0
	}

	if memo[i][j] != -1000000 {
		return memo[i][j]
	}

	res := min(dfs(triangle, i+1, j, memo), dfs(triangle, i+1, j+1, memo)) + triangle[i][j]
	memo[i][j] = res
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
