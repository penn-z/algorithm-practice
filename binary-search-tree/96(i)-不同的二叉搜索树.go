/*
 leetcode 96: 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种树。


*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n := 5
	res := numTreesV2(n)
	fmt.Println("res: ", res)
}

var memo [][]int

func numTrees(n int) int {
	memo = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}
	// 计算闭区间 [1, n] 组成的 BST 个数
	return count(1, n)
}

// 计算闭区间 [1, n] 组成的 BST 个数
// lowest: 左边界
// highest: 右边界
func count(lowest, highest int) int {
	// 当左边界大于右边界时，说明当前以当前数组最后一个元素作为根节点，该类型也属于一种情况，故return 1而非return 0
	if lowest > highest {
		return 1
	}

	if innerRes := memo[lowest][highest]; innerRes != 0 {
		return innerRes
	}

	var res int
	for i := lowest; i <= highest; i++ {
		// 以i为根节点的情况，则[lowest, i - 1]为左子树, [i+1, highest]为右子树
		leftCount := count(lowest, i-1)
		rightCount := count(i+1, highest)

		res += (leftCount * rightCount)
	}

	memo[lowest][highest] = res
	return res
}

func numTreesV2(n int) int {
	// 初始化memo
	memo := make([]int, n+1)
	return helper(n, &memo)
}

func helper(n int, memo *[]int) int {
	if n == 1 || n == 0 {
		return 1
	}

	if (*memo)[n] != 0 {
		return (*memo)[n]
	}

	count := 0
	for i := 1; i <= n; i++ {
		// 以i为根节点，则[1, i-1]为左子树，[i+1, n]为右子树
		count += helper(i-1, memo) * helper(n-i, memo)
	}

	(*memo)[n] = count
	return count
}
