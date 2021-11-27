/*
	leetcode: 503
	给定一个循环数组（最后一个元素的像一个元素是数组的第一个元素），输出每个元素的下一个更大元素。
	数字x的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出-1。

	demo1:
		input: [1, 2, 1]
		output: [2, -1, 2]
		explain: 第一个1的下一个更大的数是2；
			 	 数字2找不到下一个更大的数；
				第二个1的下一个最大的数需要循环搜索，结果也是2。

	思路： 单调栈
		[1, 2, 1] => [1, 2, 1, 1, 2]
		1. 根据题意，可以利用单调栈特性，形成一个单调递减的单调栈
		2. 原数组与 [0:len(arr)-1]拼接成一个新的数组，对此数组进行遍历
		3. 当栈不为空时，若当前元素>=栈顶元素，则不符合单调递减栈，需要弹出栈顶元素，直到当前元素<栈顶元素位置。
		4. 栈为空时，说明改元素后面没有比其更大的元素，填充-1，否则填充值为栈顶元素
			注意，因为数组是拼接的，当前结果数组索引为 i % len(originArr)
*/
package main

func main() {

}

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	if len(nums) == 0 {
		return res
	}

	lenOrigin := len(nums)
	stack := []int{}
	newNum := append(nums, nums[0:lenOrigin-1]...)
	// 逆序遍历
	for i := len(newNum) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= newNum[i] {
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			// 此时栈为空，说明该元素前面没有比起更大的元素
			res[i%lenOrigin] = -1
		} else {
			res[i%lenOrigin] = stack[len(stack)-1]
		}

		// 当前元素入栈
		stack = append(stack, newNum[i])
	}

	return res
}
