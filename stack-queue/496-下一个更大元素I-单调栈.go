/*
	leetcode: 496-下一个更大元素I
	给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
	请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
	nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1


	demo1:
		input: nums1 = [4, 1, 2], nums2 = [1, 3, 4, 2]
		output: [-1, 3, -1]
		explain:
			对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
			对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
	    	对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。

	demo2:
		input: num1 = [2, 4], num2 = [1, 2, 3, 4]
		output: [3, -1]
		explain:
			对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
    		对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。

	思路: 单调栈
		1. 根据题意，num1是num2子集，且没有重复元素，故而可以逆序遍历nums, 形成一个单调递减的栈
		2. 维护一个map，记录每个元素下一个比大的元素值，方便最后遍历nums1直接从map中读出结果

		ps: 为什么是逆序遍历nums2，因为题意求的是比当前元素大的下一个元素，栈只能倒着弹出，为了保持单调性，每次当前元素比上一个大，故这里逆序遍历nums2，可以确保需要找的数已经被处理过，正向解决问题。
*/
package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	res := []int{}

	itemMap := make(map[int]int)
	stack := []int{}
	// 逆序遍历nums2, 形成一个单调递减的单调栈
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[i] {
			// 弹出栈最后一个元素
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			// 此时单调为空，说明该元素前面没有比其大的元素
			itemMap[nums2[i]] = -1
		} else {
			// 栈不为空，说明此时栈顶元素比当前元素大
			itemMap[nums2[i]] = stack[len(stack)-1]
		}

		// 当前元素加入栈中
		stack = append(stack, nums2[i])
	}

	// 遍历num1，从map中找到结果值
	for _, num := range nums1 {
		res = append(res, itemMap[num])
	}

	return res
}
