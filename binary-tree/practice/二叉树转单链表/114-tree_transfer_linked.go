package main

import "fmt"

/*
	leetcode-114，将二叉树展开为链表:
	展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	展开后的单链表应该与二叉树 先序遍历 顺序相同。

	思路流程：
	1. 将root的左子树和右子树拉平
	2. 将root的右子树接到左子树下方
	3. 将整个左子树作为右子树
*/

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	flatten(root)
	fmt.Printf("Tree: %+v", root)
}

func flatten(root *TreeNode) {
	// 根据上述思路流程，需要先分别拉平左右子树，再进行根节点操作，故而使用后序遍历

	// base case
	if root == nil {
		return
	}

	// 左
	flatten(root.Left)

	// 右
	flatten(root.Right)

	// 根 (核心逻辑)

	// 1. 记录此时拉平后的左右子树
	left := root.Left
	right := root.Right

	// 2. 将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3. 将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}

	p.Right = right
}
