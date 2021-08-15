package main

/*
	leetcode110: 二叉平衡树
	给定一个二叉树，判断它是否是高度平衡的二叉树。
	本题中，一棵高度平衡二叉树定义为：
		一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。


	思路:
		1. 后序遍历-递归获取左右子树高度差即可
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isBalanced(root *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}

	// 出参是当前入参节点高度，若小于0则代表该节点非平衡树
	var postOrder func(*TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// 左
		heightLeft := postOrder(root.Left)
		if heightLeft < 0 {
			return -1
		}

		// 右
		heightRight := postOrder(root.Right)
		if heightRight < 0 {
			return -1
		}

		if abs(heightLeft, heightRight) > 1 {
			return -1
		}

		// 当前节点高度
		return max(heightLeft, heightRight) + 1
	}

	height := postOrder(root)
	return height >= 0
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
