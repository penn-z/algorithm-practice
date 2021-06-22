package main

import "fmt"

// tree node struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root := buildTree()
	isSymmetric := isSymmetricTree(root)
	fmt.Printf("isSymmetric:%v", isSymmetric)
}

func buildTree() (root *TreeNode) {
	node4 := &TreeNode{
		Val: 3,
	}

	node5 := &TreeNode{
		Val: 4,
	}

	node2 := &TreeNode{
		Val:   2,
		Left:  node4,
		Right: node5,
	}

	node6 := &TreeNode{
		Val: 4,
	}

	node7 := &TreeNode{
		Val: 3,
	}

	node3 := &TreeNode{
		Val:   2,
		Left:  node6,
		Right: node7,
	}

	root = &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	return
}

// 是否对称二叉树
func isSymmetricTree(root *TreeNode) (isSymmetric bool) {
	// 递归终止条件:
	// 1. 两个节点都为空，则当前子树为镜像对称
	// 2. 两个节点一个为空，当前子树不为镜像对称
	// 3. 当前两节点值不相等
	// 3. 递归比较左右子树，left.Left != right.Right 或 left.Right != right.Left, 则不为镜像对称
	var symmetircTree func(left *TreeNode, right *TreeNode) bool
	symmetircTree = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return symmetircTree(left.Left, right.Right) && symmetircTree(left.Right, right.Left)
	}

	isSymmetric = symmetircTree(root.Left, root.Right)
	return
}

// // 是否镜像对称二叉树（迭代方法）
// func isSymmetricTreeIteration(root *TreeNode) (isSymmetric bool) {
// 	//
// 	return
// }
