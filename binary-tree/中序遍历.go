package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node4 := &TreeNode{
		Val: 4,
	}

	node5 := &TreeNode{
		Val: 5,
	}

	node2 := &TreeNode{
		Val:   2,
		Right: node4,
	}

	node6 := &TreeNode{
		Val: 6,
	}

	node3 := &TreeNode{
		Val:   3,
		Left:  node5,
		Right: node6,
	}

	root := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	res := midorderTree(root)
	fmt.Println("递归中序遍历: ", res)
}

func midorderTree(root *TreeNode) (vals []int) {
	var midorder func(*TreeNode)
	midorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 左
		midorder(root.Left)

		// 中
		vals = append(vals, root.Val)

		// 右
		midorder(root.Right)
	}

	midorder(root)
	return vals
}
