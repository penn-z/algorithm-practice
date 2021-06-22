package main

import "fmt"

// tree node struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node2 := &TreeNode{
		Val: 9,
	}

	node6 := &TreeNode{
		Val: 15,
	}

	node7 := &TreeNode{
		Val: 7,
	}

	node3 := &TreeNode{
		Val:   20,
		Left:  node6,
		Right: node7,
	}

	root := &TreeNode{
		Val:   3,
		Left:  node2,
		Right: node3,
	}

	ans := maxDepth(root)
	fmt.Printf("max depth:%v", ans)

}

// 获取树最大深度
func maxDepth(root *TreeNode) (ans int32) {
	if root == nil {
		return
	}

	// 自顶向下 -- 前序遍历
	var preorder func(*TreeNode, int32)
	preorder = func(node *TreeNode, depth int32) {
		if node == nil {
			return
		}

		// 如果当前节点为叶子节点
		if node.Left == nil && node.Right == nil {
			ans = max(ans, depth)
		}

		// 左节点
		preorder(node.Left, depth+1)
		// 右节点
		preorder(node.Right, depth+1)
	}

	preorder(root, 1)
	return
}

// 取大值
func max(x, y int32) int32 {
	if x > y {
		return x
	}

	return y
}
