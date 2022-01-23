/*
	leetcode 95: 不同的二叉搜索树II
	给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

	demo1:
		输入：n = 3
		输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

	思路: 递归
	1. 穷举root节点的所有可能
	2. 递归构造出左右子树的所有合法BST
	3. 给root节点穷举所有左右子树的组合

*/
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return build(1, n)
}

func build(lower, higher int) []*TreeNode {
	if lower > higher {
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}

	// 1. 穷举root节点所有可能
	for i := lower; i <= higher; i++ {
		// 2. 递归构造出所有合法的BST
		leftTree := build(lower, i-1)
		rightTree := build(i+1, higher)
		// 3. 给root节点穷举所有左右子树的组合
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTree {
			for _, right := range rightTree {
				curTree := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}

				res = append(res, curTree)
			}
		}
	}

	return res
}
