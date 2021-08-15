/*
	leetcode 538:
		给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

	思路:
		1. 题意求新树，当前节点的新值为大于等于当前节点的节点值之和
		2. 可以利用中序遍历数组升序特性，但是该题是需要求大于等于当前节点其他节点的值之和，可把中序遍历倒过来，即是降序数组，即符合题意
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func main() {

}

func convertBST(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	// 右
	traverse(root.Right)

	// 中
	sum += root.Val
	root.Val = sum

	// 左
	traverse(root.Left)
	return
}
