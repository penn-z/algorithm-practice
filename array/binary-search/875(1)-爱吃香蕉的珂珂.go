/*
	珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
	珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
	珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
	返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。


	二分搜索模板无非就是关于target, x(自变量), f(x)的单调函数

	思路：
		运用二分搜素模板
		1. 确认x, target, f(x)分别是什么: x自变量为每小时吃香蕉速度K, f(x)则是吃完所有香蕉所需的时间，target则是H，是f(x)的最大约束
		2. 写出f(x)函数表达式
		3. 找到x的取值范围作为二分查找的搜索区间，初始化left和right变量
		4. 根据题意，可知x最小是1， 最大值为piles数组中元素最大值
		5. 确定是左侧还是右侧的二分搜索，画出f(x)伪单调函数，求x最小值，可知是求左侧边界
		6. f(x)是关于x的单调递减函数，相当于降序数组，二分查找对半操作时要与升序数组相反，注意
*/
package main

import "fmt"

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	// piles := []int{30, 11, 23, 4, 20}
	// h := 5
	res := minEatingSpeed(piles, h)
	fmt.Printf("===res:%d\n", res)
}

func minEatingSpeed(piles []int, h int) int {
	//  base case
	if len(piles) == 0 {
		return -1
	}

	// 确定左右边界left, right
	left := 1

	// 求出piles中元素最大值
	right := getMaxValue(piles)

	for left <= right {
		mid := (left + right) / 2
		if f(piles, mid) == h {
			// 不急于返回mid，求左侧边界，则收缩右侧边界
			right = mid - 1
		} else if f(piles, mid) < h {
			// 需要让f(x)返回的值更大一些，由于是单调减函数，故而需要收缩右边界，才能使f(x)增大
			right = mid - 1
		} else if f(piles, mid) > h {
			// 需要让f(x)返回的值更小一些，由于单调减函数，故而需要收缩左边界，才能使f(x)减小
			left = mid + 1
		}
	}

	// for退出条件是left = right + 1，但此处不存在越界问题
	return left

}

// 求f(x)关于x的单调函数，出参为耗时h
// x即为每小时吃香蕉速度
func f(piles []int, x int) int {
	hours := 0
	for _, num := range piles {
		hours += num / x
		// 当前整数小时吃不完，需要额外一小时吃
		if num%x > 0 {
			hours++
		}
	}

	return hours
}

// 求出piles中元素最大值
func getMaxValue(piles []int) int {
	max := 0
	for _, num := range piles {
		if num > max {
			max = num
		}
	}

	return max
}
