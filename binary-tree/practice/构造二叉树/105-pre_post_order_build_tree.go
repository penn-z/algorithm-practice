package main

import "fmt"

// 从前序与中序遍历序列中构造二叉树, leetcode-105

/*
	思路:
		1. 前序遍历第一个元素即为根节点
		2. 中序遍历根节点位置可将中序遍历数组分为左右子树
		3. 定义一个hashMap保存中序遍历中，元素与索引的位置关系，因为从前序序列中拿到根节点后，要在中序序列中查找对应的位置，从而将数组分为左右子树
		4. 根据根节点的位置确定左子树和右子树在中序数组和前序数组中的左右边界位置
		5. 递归构造左子树和右子树
		6. 返回根节点结束
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	inorderIndexMap = make(map[int]int)
	pre             = []int{}
)

func main() {
	inorder := []int{4, 2, 8, 5, 9, 1, 6, 10, 3, 7}
	preorder := []int{1, 2, 4, 5, 8, 9, 3, 6, 10, 7}
	// inorder := []int{3, 9, 20, 15, 7}
	// preorder := []int{9, 3, 15, 20, 7}
	root := buildTree(inorder, preorder)
	fmt.Printf("tree %+v", root)
}

// 构造树
func buildTree(inorder []int, preorder []int) *TreeNode {
	for index, val := range inorder {
		inorderIndexMap[val] = index
	}

	pre = preorder

	return buildTreeRecusion(0, len(inorder)-1, 0, len(preorder)-1)
}

// 递归构造树
func buildTreeRecusion(inorderStart, inorderEnd, preorderStart, preorderEnd int) *TreeNode {
	if inorderEnd < inorderStart || preorderEnd < preorderStart {
		return nil
	}

	// if preorderEnd == preorderStart {
	// 	return nil
	// }

	// 1. 根据前序遍历结果，获取根节点
	rootVal := pre[preorderStart]
	// 2. 获取根节点对应索引
	rootIndex := inorderIndexMap[rootVal]
	fmt.Printf("===rootVal:%d, rootIndex:%d\n", rootVal, rootIndex)

	node := &TreeNode{
		Val: rootVal,
	}

	node.Left = buildTreeRecusion(inorderStart, rootIndex-1, preorderStart+1, preorderStart+rootIndex-inorderStart)
	node.Right = buildTreeRecusion(rootIndex+1, inorderEnd, preorderStart+rootIndex-inorderStart+1, preorderEnd)
	return node
}
