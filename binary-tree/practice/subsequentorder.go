package main

import "fmt"

// tree node struct
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

	vals := subsequentorderTravesal(root)

	fmt.Printf("vals:%v\n", vals)
}

func subsequentorderTravesal(root *TreeNode) (vals []int) {
	var subsequentorder func(*TreeNode)

	subsequentorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 左子结点
		subsequentorder(node.Left)

		// 右子结点
		subsequentorder(node.Right)

		// 根节点
		vals = append(vals, node.Val)

	}

	subsequentorder(root)
	return
}
