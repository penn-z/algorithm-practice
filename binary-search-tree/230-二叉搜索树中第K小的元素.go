package main

/*
	给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）

	思路：
		1. 利用二叉搜索树中序遍历升序的特性，遍历至第k个即是第k个最小元素
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	seq int
	res int
)

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	// base case
	if root == nil {
		return 0
	}

	// 中序遍历
	seq = k
	midOrder(root)

	return res
}

func midOrder(root *TreeNode) {
	if root == nil {
		return
	}

	// 中序遍历
	// 左
	midOrder(root.Left)

	// 中
	seq--
	if seq == 0 {
		res = root.Val
		return
	}

	// 右
	midOrder(root.Right)

	return
}
