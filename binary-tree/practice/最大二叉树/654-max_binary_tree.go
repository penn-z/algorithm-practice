package main

/*
	leetcode-654 最大二叉树

	给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
		二叉树的根是数组 nums 中的最大元素。
		左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
		右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
		返回有给定数组 nums 构建的 最大二叉树 。


	思路：
		1. 关键找出nums数组中的最大值max, 则最大值元素 左边的元素为左子树，右边元素为右子树
		2. 构造以此max元素为root的树
		3. 依次对左右子树进行递归构造操作
		4. 可知为前序遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// base case
	if len(nums) < 1 {
		return nil
	}

	// 找出数组中最大值元素
	max := ^int(^uint(0) >> 1)
	maxIndex := -1
	for i, item := range nums {
		if item > max {
			max = item
			maxIndex = i
		}
	}

	root := &TreeNode{
		Val: max,
	}

	// 左子树递归
	root.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	// 右子树递归
	return root
}
