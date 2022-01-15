/*
	leetcode 135 - 分发糖果
	n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

	你需要按照以下要求，给这些孩子分发糖果：
		* 每个孩子至少分配到 1 个糖果。
		* 相邻两个孩子评分更高的孩子会获得更多的糖果。
	请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

	demo1:
		输入：ratings = [1,0,2]
		输出：5
		解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。

	demo2:
		输入：ratings = [1,2,2]
		输出：4
		解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
			第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。


	思路： 贪心算法，求出局部最优，推出全局最优。

	左规则: 从左向右遍历，先确定右边评分大于左边的情况，
		局部最优：只要右边评分比左边大，则右孩子多一个糖果
		全局最优：相邻的孩子中，评分高的右孩子比左孩子获得更多的糖果
		即ratings[i] > ratings[i-1]， 那么[i]的糖果一定比[i-1]多一个，所以贪心: candyVec[i] = candyVec[i-1] + 1

	右规则：从右往左遍历，确定左边评分大于右边的情况
		局部最优：只要左边评分比右边大，左孩子要保持比右孩子大，则 左边小与右边时， candyVec[i] = candyVec[i+1] + 1
		全局最优：相邻的孩子中，评分高的孩子获得更多的糖果
		ratings[i] > ratings[i+1]， 左孩子要保持比右孩子大， 那么 candyVec[i] = max(candyVec[i], candyVec[i-1] + 1)

*/
package main

func main() {

}

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	// 一个数组记录每个孩子可分得糖果数量
	candyArr := make([]int, len(ratings))
	// 初始值，每人至少1个糖果
	for i := range candyArr {
		candyArr[i] = 1
	}

	// 左规则: 从左往右遍历，若 rating[i] > rating[i-1]，则i号孩子要比i-1号孩子糖果多1
	// 右规则: 从右往左遍历，若 rating[i] > rating[i+1]，则i号孩子要比i+1号孩子糖果多1

	candyArr[0] = 1
	for i := 1; i < len(ratings); i++ {
		// ratings[i] > ratings[i-1]， 那么[i]的糖果一定比[i-1]多一个，所以贪心: candyVec[i] = candyVec[i-1] + 1
		if ratings[i] > ratings[i-1] {
			candyArr[i] = candyArr[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		// ratings[i] > ratings[i+1]， 左孩子要保持比右孩子大， 那么 candyVec[i] = max(candyVec[i], candyVec[i-1] + 1)
		if ratings[i] > ratings[i+1] {
			if candyArr[i] <= candyArr[i+1] {
				candyArr[i] = candyArr[i+1] + 1
			}
		}
	}

	res := 0
	for i := range candyArr {
		res += candyArr[i]
	}

	return res
}
