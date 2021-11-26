/*
	请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。

	demo1:
		input: temperatures = [73,74,75,71,69,72,76,73]
		output: [1,1,4,2,1,1,0,0]

	demo2:
		input: temperatures = [30,40,50,60]
		output: [1,1,1,0]

	demo3:
		input: temperatures = [30,60,90]
		output: [1,1,0]

	思路: 单调栈
		1. 根据题意，可以利用单调栈，逆序遍历，形成一个递减的单调栈
		2. 当前元素的下标入栈，如果栈顶元素下标对应的原数组值比当前数值小，重复该过程，直到栈顶元素索引对应值比当前元素值大为止
		3. 当前元素索引入栈

*/
package main

import "fmt"

func main() {
	nums := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(nums)
	fmt.Println("======res:", res)
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures), len(temperatures))
	if len(temperatures) == 0 {
		return res
	}

	stack := []int{}

	// 逆序遍历
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			// 弹出栈顶元素，不满足单调性
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			// 此时单调栈为空，说明该元素前面没有比其大的元素
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1] - i
		}

		// 索引入栈
		stack = append(stack, i)
	}

	return res
}
