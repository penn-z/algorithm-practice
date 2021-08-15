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

	var postOrder func(*TreeNode) (int, bool)

	postOrder = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}

		// 左
		heightLeft, balance := postOrder(root.Left)
		if !balance {
			return 0, false
		}

		// 右
		heightRight, balance := postOrder(root.Right)
		if !balance {
			return 0, false
		}

		// 当前节点高度
		var curNodeHeight int
		if heightLeft > heightRight {
			curNodeHeight = heightLeft + 1
		} else {
			curNodeHeight = heightRight + 1
		}

		var isBalance bool
		if abs(heightLeft, heightRight) <= 1 {
			isBalance = true
		}

		return curNodeHeight, isBalance
	}

	_, isBlance := postOrder(root)
	return isBlance
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}
