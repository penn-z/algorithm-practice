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

func isValidBSTV3(root *TreeNode) bool {
	return isValidBstTraverse(root, nil, nil)
}

/*
	思路:
		1. 根据BST的特性。左边的小于根节点的值，右边的大于根节点的值。并且对于每一棵子树都是如此。
		2. 所以我们可以直接递归的对左右子树的值与根节点的值进行比较。
		3. 左子树的值小于当前根节点的值，将当前根节点的值作为最大值传入左子树，左子树的值都小于他，递归处理；右子树的值都大于根节点的值，将根节点的值作为最小值传入右子树，右子树的值都大于他。
*/
func isValidBstTraverse(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}

	// 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= max.Val {
		return false
	}

	// 限制左子树的最大值是 root.val, 右子树的最小值是 root.val，递归处理检查左右子树
	return isValidBstTraverse(root.Left, min, root) && isValidBstTraverse(root.Right, root, max)
}

/*
	判断是否二叉搜索树模板: INT_MIN子树中最小值, INT_MAX子树中最大值
	is_BST(root, INT_MIN, INT_MAX)
*/
