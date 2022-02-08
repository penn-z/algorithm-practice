/*
	leetcode 104 - 二叉树的最大深度
	给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	前序遍历 - 自顶向下
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	// 定义前序遍历递归函数
	var preorder func(*TreeNode, int)
	preorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		// 中
		// 如果当前节点为叶子节点，则需要比较当前叶子节点层级大小与答案大小，取大值
		if node.Left == nil && node.Right == nil {
			// 更新ans
			ans = max(ans, depth)
		}

		// 左节点
		preorder(node.Left, depth+1)

		// 右节点
		preorder(node.Right, depth+1)
	}

	preorder(root, 1)
	return ans
}

var (
	// 记录最大深度
	res = 0
	// 记录遍历到节点的深度
	curDepth = 0
)

/*
	回溯思路
*/
func maxDepthV2(root *TreeNode) int {
	res, curDepth = 0, 0
	traverse(root)
	return res
}

// 回溯遍历
func traverse(root *TreeNode) {
	if root == nil {
		// 叶子节点，更新最大深度
		res = max(res, curDepth)
		return
	}

	// 前序位置，进入一个节点时，需要增加curDepth
	curDepth++
	traverse(root.Left)
	traverse(root.Right)
	// 后序位置，离开一个节点后，需要减小curDepth
	curDepth--
}

// 分解问题
func maxDepthV3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 分别计算左右子树最大深度
	leftMax := maxDepthV3(root.Left)
	rightMax := maxDepthV3(root.Right)
	// 整棵树的最大深度等于左右子树的最大深度取最大值
	// 然后再加上根节点自身
	res := max(leftMax, rightMax) + 1
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
