package main

import "fmt"

/*
	leetcode-652 给定一个二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根节点即可。
	两棵树重复是指它们具有相同的结构以及相同的结点值。

	思路:
		1. 当前结点的子树是什么？ （知道子树后再知道自身，可知为后序遍历）
		2. 其他结点的子树是什么？
		3. 需要map记录每个结点子树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var treeMap = make(map[string]int)
var resList = []*TreeNode{}

func main() {

}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	findSubTree(root)
	return resList
}

func findSubTree(root *TreeNode) string {
	// base case
	if root == nil {
		// nil结点用 # 表示
		return "#"
	}

	// 后序遍历

	// 左
	left := findSubTree(root.Left)

	// 右
	right := findSubTree(root.Right)

	// 根，以当前节点为根节点的树字符串
	subTree := fmt.Sprintf("%s,%s,%d", left, right, root.Val)

	if treeMap[subTree] == 1 {
		// 出现过1次，则此结点符合题意
		resList = append(resList, root)
	}

	treeMap[subTree] = treeMap[subTree] + 1
	return subTree
}
