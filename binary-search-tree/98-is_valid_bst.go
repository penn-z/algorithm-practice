package main

import "fmt"

/*
	leetcode-98 验证二叉搜索树
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。

	假设一个二叉搜索树具有如下特征:
		1. 节点的左子树只包含小于当前节点的数。
		2. 节点的右子树只包含大于当前节点的数。
		3. 所有左子树和右子树自身必须也是二叉搜索树。

	思路:
		1. 二叉搜索树是中序遍历结果，是个顺序数组
		2. 构建二叉搜索树中序遍历结果数组
		3. check数组是否顺序数组
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nodeValList = []int{}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	ret := isValidBST(root)
	fmt.Printf("===ret:%v", ret)
}

func isValidBST(root *TreeNode) bool {
	nodeValList = []int{}
	midOrder(root)
	// check is order slice
	lenNodeList := len(nodeValList)
	for i := 0; i < lenNodeList-1; i++ {
		if nodeValList[i] >= nodeValList[i+1] {
			return false
		}
	}

	return true
}

func midOrder(root *TreeNode) {
	// base
	if root == nil {
		return
	}

	// 左
	midOrder(root.Left)
	// 中
	nodeValList = append(nodeValList, root.Val)
	// 右
	midOrder(root.Right)
	return
}
