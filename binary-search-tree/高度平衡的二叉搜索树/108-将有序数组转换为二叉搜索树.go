package main

import "fmt"

/*
	leetcode108: 将有序数组转换为二叉搜索树

	给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

	高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。


	输入：nums = [-10,-3,0,5,9]
	输出：[0,-3,9,-10,null,5]
	解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：


	思路:
		1. 题意给出的是升序数组，可选择下标为overflow(n/2)=mid作为根节点，则[0, mid-1]为左子树，[mid+1, len(arrs)-1]为右子树
		2. 利用中序遍历特性，递归左右子树构造为二叉搜索树即可

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test := []int{1, 2, 3, 4, 5}
	mid := len(test) / 2

	fmt.Println(mid)
	fmt.Println(test[0:mid])
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	// 中间位置
	mid := len(nums) / 2

	// 中序遍历
	root := &TreeNode{
		Val: nums[mid],
	}

	// 左
	root.Left = sortedArrayToBST(nums[0:mid]) // 0:mid, 不包含mid

	// 右
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
