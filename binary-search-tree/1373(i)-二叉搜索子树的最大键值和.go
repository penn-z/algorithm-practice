package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var res int

func maxSumBST(root *TreeNode) int {
	res = 0
	maxSumBSTTraverse(root)
	return res
}

func maxSumBSTTraverse(root *TreeNode) {
	// 1. 是否BST
	if isBST(root, nil, nil) {
		// 2. 算出当前节点为BST的子树值和
		sumNodeValue(root)
		return
	}

	// 不是BST，递归寻找子节点
	maxSumBSTTraverse(root.Left)
	maxSumBSTTraverse(root.Right)
	return
}

func isBST(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	// base case
	if root != nil {
		return true
	}

	// 若root.Val，不符合min和max的限制，则不是合法BST
	if min != nil && root.Val < min.Val {
		return false
	}

	if max != nil && root.Val > max.Val {
		return false
	}

	return isBST(root.Left, min, root) && isBST(root.Right, root, max)
}

// res结果集
func sumNodeValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 二叉搜索树节点和的最大值，一定在子节点中出现
	sum := root.Val + sumNodeValue(root.Left) + sumNodeValue(root.Right)
	// 求出sum, 与res较大值
	res = getGreaterValue(res, sum)
	return sum
}

func getGreaterValue(x, y int) int {
	if x > y {
		return x
	}

	return y
}
