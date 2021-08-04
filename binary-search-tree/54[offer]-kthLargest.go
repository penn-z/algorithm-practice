package main

/*
	给定一棵二叉搜索树，请找出其中第k大的节点。
	示例:
		输入: root = [3,1,4,null,2], k = 1
		输入: 4


		输入: root = [5,3,6,2,4,null,null,1], k = 3
		输出: 4


	解题思路:
		二叉搜索树中序遍历是升序，题意要求第k大。则可以按照中序遍历操作倒序，即是降序数组

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var globalK int // 记录当前操作顺位
var res int     // 结果

func kthLargest(root *TreeNode, k int) int {
	// base case
	if root == nil {
		return 0
	}

	globalK = k
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	// base case
	if root == nil {
		return
	}

	// 右
	dfs(root.Right)

	globalK -= 1

	// 中
	if globalK == 0 {
		res = root.Val
		return
	}

	// 左
	dfs(root.Left)
	return
}
