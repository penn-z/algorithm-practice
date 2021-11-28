/*
	leetcode 42: 接雨水
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

	demo1:
		输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
		输出：6
		解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。


	思路1：单调栈
		可接雨水的区域为低洼处，则高度必须先减，然后上升，才能形成低洼区域。故而可以单调递减栈维护柱子高度的索引，遍历柱子数组，遇到递减入栈，否则弹栈算低洼区域面积，累加到答案中。
		1. 用栈来保存数组索引下标
		2. 从左到右遍历柱子数组, height[i]为当前柱子:
			a. 当栈非空且height[i] > stack.top()
				- 说明此处会形成低洼区域，可被弹出，top = stack.top()
				- 计算积水宽度：即当前元素与新栈顶元素的距离，准备进行填充操作。 distance = i - st.top() - 1
				- 计算积水高度: bound_height = min(height[i], st.top()) - top   (注意: st.top()为新栈顶元素，top才是刚才弹出的栈顶元素)
				- 累加积水量: ans += distance * bound_height
			b. 将当前柱子索引入栈(经过a的循环，到此处是满足单调递减栈的)
			c. i继续移动

	思路2: 暴力解法
		对于数组中的每个元素，找出水能达到的最高位置，等于两边最大高度的较小值减去当前高度值。


	思路3: 暴力解法+备忘录优化 = 动态规划
		1. 对于思路2， 每次对需要求出height[i]左右柱子left_max, right_max最大值，可以考虑用map存储下来结果，方便后续直接使用


*/

package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	resV2 := trapV2(height)
	fmt.Println("resV2: ", resV2)
}

func trap(height []int) int {
	var res int
	if len(height) == 0 {
		return res
	}

	stack := []int{}
	for i := 0; i < len(height); i++ {
		// 当前栈不为空，且当前柱子高度 > stack.top()
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]

			// 弹出栈顶元素
			stack = stack[:len(stack)-1]

			// 栈已经清空，左边没有柱子了，无法积水，退出
			if len(stack) == 0 {
				break
			}

			// 计算宽度
			left := stack[len(stack)-1]
			distance := i - left - 1
			// 计算高度
			boundHeight := minNum(height[i], height[left]) - height[top]

			res += distance * boundHeight
		}

		// 当前柱子索引入栈
		stack = append(stack, i)
	}

	return res

}

func minNum(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// 思路2：暴力解法
func trapV2(height []int) int {
	var res int
	if len(height) == 0 {
		return res
	}

	for i, _ := range height {
		// 初始化左右高度的最大值left_max, right_max
		var left_max, right_max int = 0, 0
		// 左边柱子最大值
		for j := i; j >= 0; j-- {
			left_max = maxNum(left_max, height[j])
		}

		// 右边柱子最大值
		for j := i; j <= len(height)-1; j++ {
			right_max = maxNum(right_max, height[j])
		}

		// 当前积水量加入结果中
		res += minNum(left_max, right_max) - height[i]
	}

	return res
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// 思路3: 暴力解法+备忘录优化 = 动态规划
func trapV3(height []int) int {
	var res int
	if len(height) == 0 {
		return res
	}

	height_len := len(height)
	// 初始化left_max, right_max数组
	left_max_arr := make([]int, height_len, height_len)
	right_max_arr := make([]int, height_len, height_len)

	left_max_arr[0] = height[0]
	for i := 1; i < height_len; i++ {
		left_max_arr[i] = maxNum(left_max_arr[i-1], height[i])
	}

	right_max_arr[height_len-1] = height[height_len-1]
	for i := height_len - 2; i >= 0; i-- {
		right_max_arr[i] = maxNum(right_max_arr[i+1], height[i])
	}

	// 遍历height，求积水量
	for i, _ := range height {
		res += minNum(left_max_arr[i], right_max_arr[i]) - height[i]
	}

	return res
}
