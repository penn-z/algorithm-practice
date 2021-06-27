package main

import "fmt"

// 从中序与后序遍历中构造二叉树, leetcode-106

/*
	思路:
		1. 后序遍历中最后一个元素即为根节点
		2. 中序遍历中根节点位置可将中序遍历数组分为左右子树
		3. 定义一个hashMap保存中序遍历中，元素和索引的位置关系，因为从后序序列中拿到根节点后，要在中序序列中查找对应的位置，从而将数组分为左子树和右子树
		4. 根据根节点的位置确定左子树和右子树在中序数组和后续数组中的左右边界位置
		5. 递归构造左子树和右子树
		6. 返回根节点结束
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义map记录中序遍历中，元素和索引的位置关系
var inorderIndexMap = make(map[int]int)
var post = []int{}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Printf("tree: %+v", root)

}

// 构造树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 思路:

	// 定义map记录中序遍历中，元素和索引的位置关系
	// inorderIndexMap := make(map[int]int)
	for index, val := range inorder {
		inorderIndexMap[val] = index
	}

	post = postorder

	root := buildTreeRecursion(0, len(inorder)-1, 0, len(postorder)-1)
	return root
}

// 递归构造树
func buildTreeRecursion(inorderStart, inorderEnd, postorderStart, postorderEnd int) *TreeNode {
	// 越界，直接退出
	if inorderEnd < inorderStart || postorderEnd < postorderStart {
		return nil
	}

	// 1. 根据后序遍历结果，获取根节点
	rootVal := post[postorderEnd]
	// 2. 获取对应索引
	rootIndex := inorderIndexMap[rootVal]

	node := &TreeNode{
		Val: rootVal,
	}

	node.Left = buildTreeRecursion(inorderStart, rootIndex-1, postorderStart, postorderStart+rootIndex-inorderStart-1)
	node.Right = buildTreeRecursion(rootIndex+1, inorderEnd, postorderStart+rootIndex-inorderStart, postorderEnd-1)

	return node
}
