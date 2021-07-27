package main

import (
	"fmt"
	"math"
)

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

// 是否搜索二叉树
/*
	思路:
		1. 迭代中序遍历树
		2. 迭代过程中实时检查当前节点的val是否大于前一个中序遍历到的节点的之即可
*/
func isValidBSTV2(root *TreeNode) bool {
	// 先将根节点root和所有左孩子入栈
	stack := []*TreeNode{}

	tmpNodeVal := math.MinInt64
	for len(stack) > 0 || root != nil {
		// 遍历左孩子，加入栈中，直至root为空
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 处理中间结点
		// 弹出栈顶元素，将此元素加入结果中，到达它的右孩子，并将这个节点当做cur重新执行步骤，直至栈为空
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 中间节点
		if top.Val <= tmpNodeVal {
			return false
		}

		tmpNodeVal = top.Val
		root = top.Right
	}

	return true
}
