package main

import (
	"fmt"
)

// 路径和
// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

// tree node struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node4 := &TreeNode{
		Val: 4,
	}

	node5 := &TreeNode{
		Val: 5,
	}

	node2 := &TreeNode{
		Val:   2,
		Right: node4,
	}

	node6 := &TreeNode{
		Val: 6,
	}

	node3 := &TreeNode{
		Val:   3,
		Left:  node5,
		Right: node6,
	}

	root := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	// isExists := hasPathSumIteration(root, 9) // 迭代-深度优先
	// fmt.Printf("isExsits:%v\n", isExists)

	// isExists := hasPathSumIterationBFS(root, 9) // 迭代-广度优先
	isExists := hasPathSumRecursion(root, 10) // 递归
	fmt.Printf("root isExists:%v\n", isExists)
}

// 迭代(累加)
func hasPathSumIteration(root *TreeNode, targetSum int) bool {
	// 迭代法，累加
	// DFS，栈维护，先进后出，先右子树入栈，后左子树入栈

	if root == nil {
		return false
	}

	// 初始化栈
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// top.Val = sum + top.Val
		// 累加至叶子节点时，若结果等于sum，则该路径符合题意
		if top.Left == nil && top.Right == nil {
			if top.Val == targetSum {
				return true
			}
		}

		// 右子树累减，然后入栈
		if top.Right != nil {
			top.Right.Val += top.Val
			stack = append(stack, top.Right)
		}

		// 左子树累减，然后入栈
		if top.Left != nil {
			top.Left.Val += top.Val
			stack = append(stack, top.Left)
		}

	}

	return false
}

// 迭代（累加）
func hasPathSumIterationBFS(root *TreeNode, targetSum int) bool {
	// 迭代法，累加
	// BFS, 队列，左子树先入栈，后右子树入栈，先进先出
	if root == nil {
		return false
	}

	// 初始化队列
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top.Left == nil && top.Right == nil {
			// 叶子节点
			if top.Val == targetSum {
				return true
			}
		}

		// 左子树入栈
		if top.Left != nil {
			top.Left.Val += top.Val
			queue = append(queue, top.Left)
		}

		// 右子树入栈
		if top.Right != nil {
			top.Right.Val += top.Val
			queue = append(queue, top.Right)
		}
	}

	return false
}

// 递归
func hasPathSumRecursion(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var pathSumRecusion func(*TreeNode, int) bool
	pathSumRecusion = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}

		// 该点为叶子节点时
		if node.Left == nil && node.Right == nil {
			return node.Val == sum
		}

		return pathSumRecusion(node.Left, sum-node.Val) || pathSumRecusion(node.Right, sum-node.Val)
	}

	return pathSumRecusion(root, targetSum)
}
