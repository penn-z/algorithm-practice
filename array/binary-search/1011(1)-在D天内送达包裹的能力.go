/*
	传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。
	传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
	返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

	思路:
		1. 分析题意，确定x, target, f(x)分别是什么
		2. x即时所求的最低运力，f(x)则是在该运力下消耗的天数，target则是约束f(x)的上下限
		3. 可知是单调递减函数，写出x, f(x)关系的具体实现
		4. 可知x的最小值1，最大值是所有包裹重量之和
*/

package main

import "fmt"

func main() {
	// arrs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// days := 5
	arrs := []int{1, 2, 3, 1, 1}
	days := 4
	res := shipWithinDays(arrs, days)
	fmt.Printf("res:%d\n", res)
}

func shipWithinDays(weights []int, days int) int {
	// base case
	if len(weights) == 0 {
		return -1
	}

	// 最小最大值为
	left, right := maxTransportNums(weights)

	fmt.Printf("min:%d, max: %d\n", left, right)

	for left <= right {
		mid := (left + right) / 2

		fmt.Println(">>>>>>>>>>>>>>>>>>")
		fmt.Printf("left:%d\n", left)
		fmt.Printf("right:%d\n", right)
		fmt.Printf("mid:%d\n", mid)
		// fmt.Printf("days:%d\n", days)
		fx := f(weights, mid)
		fmt.Printf("fx:%d\n", fx)
		fmt.Println("<<<<<<<<<<<<<<<<<<")
		if fx == days {
			// 不急于返回，搜索最小左边界，则收缩右边界
			right = mid - 1
			// if mid ==
		} else if fx < days {
			// 需要让f(x)更大，f(x)是单调减函数，则需要缩小x，则需要收缩右侧边界
			right = mid - 1
		} else if fx > days {
			// 需要让f(x)更小，f(x)是单调减函数，则需要增大x，则需要收缩左侧边界
			left = mid + 1
		}
	}

	// 终止条件是left = right + 1，不会发生越界行为
	return left
}

// fx出参为需要消耗的累计天数
// f(x)随着x的增加而单调递减
func f(weights []int, x int) int {
	calcDays := 0 // 需要运输的天数
	for i := 0; i < len(weights); {
		cap := x // 每天运载力
		for i < len(weights) {
			if cap < weights[i] {
				// 超过当前运载力
				break
			} else {
				cap -= weights[i]
			}

			i++
		}

		calcDays++
	}

	return calcDays
}

// 获取最大的运载能力
func maxTransportNums(weights []int) (min, max int) {
	for _, w := range weights {
		max += w
		if min < w {
			min = w
		}
	}

	return
}
