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


*/

package main

func main() {

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
