/*
	给定一个已按照 非递减顺序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
	函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。
	你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。

	demo1:
		输入: numbers = [2, 7, 11, 15], target = 9
		输出: [1, 2]

	demo2:
		输入: numbers = [2, 3, 4], target = 6
		输出: [1, 3]

	demo3:
		输入: numbers = [-1, 0], target = -1
		输出: [1, 2]


	思路:
		因为是非递减数组，故而可以使用快慢指针比较
		1. 使用头尾双指针，不断向中间移动
		2. 当头指针大于大于尾指针时，即退出循环
*/

package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	res1 := twoSum(numbers, 9)
	fmt.Printf("===res1:%d\n", res1)

	nubmers2 := []int{2, 3, 4}
	res2 := twoSum(nubmers2, 6)
	fmt.Printf("===res2:%d\n", res2)
}

// twoSum
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum > target {
			// 两数之和大于target，则缩小两数之和，尾指针向左移动
			j--
		} else if sum < target {
			// 两数之和小于target，则扩大两数之和，头指针向右移动
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}

	return nil
}
