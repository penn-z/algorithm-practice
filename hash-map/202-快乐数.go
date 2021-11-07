/*
	leetcode 202: 快乐数
	编写一个算法来判断一个数 n 是不是快乐数。

	「快乐数」定义为：
		对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
		然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
		如果 可以变为  1，那么这个数就是快乐数。
		如果 n 是快乐数就返回 true ；不是，则返回 false 。

	思路:
		1. 求出各数位平方之和
		2. 定义set记录每个n是否重复出现
		3. 遍历入参n, 不断在循环中求解数位平方之和
*/
package main

import "fmt"

func main() {
	n := 19
	res := isHappy(n)
	fmt.Printf("res: %+v", res)
}

func isHappy(n int) bool {
	// 根据题意，可能出现重复，则用set来记录每次sum记录
	set := map[int]bool{}
	// 无限循环
	for {
		//获取数位平方和
		sum := getSum(n)
		// 符合快乐数原则，返回结果即可
		if sum == 1 {
			return true
		}

		// 如果这个sum曾经出现过，说明陷入无限循环，返回false即可
		if _, has := set[n]; has {
			return false
		} else {
			set[n] = true
		}

		n = sum // 当前数字替换为数位平方和数字
	}
}

// 求数位平方和
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10) // 最最低位平方和
		n = n / 10                 // 取得最低位平方和后，去掉最低位
	}

	return sum
}
