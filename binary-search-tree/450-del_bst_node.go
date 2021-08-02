package main

/*
	给定一个二叉搜索树的根节点root和一个值key, 删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。
	返回二叉搜索树（有可能被更新）的根节点和引用。

	两个步骤：
		1. 首先找到需要删除的节点；
		2. 如果找到了，删除它
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// BST遍历框架:
/*

func BST(root *TreeNode, target int) {
	if root.Val == target {
		// 找到目标，做点什么
	}

	if root.Val < target {
		BST(root.Right, target)
	}

	if root.Val > target {
		BST(root.Left, target)
	}
}

*/

func deleteNodeBstRecusion(root *TreeNode, target int) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// 找到目标，做点什么
	if root.Val == target {
		// 情况1，当前节点没有子节点
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// 情况2， 当前节点仅有一个子节点
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		// 情况3， 当前节点有两个子节点
		// 找到左子树中最大的节点或者右子树中最小节点
		// 这里拿右子树中最小节点示例
		minNode := getMin(root.Right) // 获取右子树最小节点

		// 替换当前节点值
		root.Val = minNode.Val
		// 递归删除右子树中最小节点
		root.Right = deleteNodeBstRecusion(root.Right, minNode.Val)
	} else if root.Val > target {
		// 当前节点大于目标节点，左子树中操作
		root.Left = deleteNodeBstRecusion(root.Left, target)
	} else if root.Val < target {
		// 当前节点小于目标节点，右子树中操作
		root.Right = deleteNodeBstRecusion(root.Right, target)
	}

	return root
}

func getMin(root *TreeNode) *TreeNode {
	// BST最小节点即是最左边节点
	for root.Left != nil {
		root = root.Left
	}

	return root
}
