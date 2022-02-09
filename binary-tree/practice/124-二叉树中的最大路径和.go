/*
	leetcode 124 - 二叉树中的最大路径和
	路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
	路径和 是路径中各节点值的总和。
	给你一个二叉树的根节点 root ，返回其 最大路径和 。


	输入：root = [1,2,3]
	输出：6
	解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6


	输入：root = [-10,9,20,null,null,15,7]
	输出：42
	解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42


	思路1: 分解问题 - 后序遍历递归
		1. 分解问题：
			1. 当前节点的最大路径和等于左节点的路径和 + 右节点路径和
*/
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最大路径和
var (
	maxInt        = int(^uint(0) >> 1)
	minInt        = ^int(maxInt)
	maxPathSumNum = minInt
)

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxInt := int(^uint(0) >> 1)
	minInt := ^int(maxInt)
	maxPathSumNum = minInt
	// 计算单边路径和时顺边计算最大路径和
	getOneSideMax(root)

	return maxPathSumNum
}

// 计算以根节点root为起点的最大单边路径和
func getOneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左子树路径总和与0比较，如果小于0，则不计入总路径中
	leftMaxSum := max(0, getOneSideMax(root.Left))
	// 右子树路径总和与0比较，如果小于0，则不计入总路径中
	rightMaxSum := max(0, getOneSideMax(root.Right))
	// 后序遍历位置，计算出当前节点的路径和
	curSum := leftMaxSum + rightMaxSum + root.Val

	// 求出当前节点的最大路径和，并更新最大路径和
	maxPathSumNum = max(maxPathSumNum, curSum)

	// 实现函数定义，左右子树的最大单边路径和加上根节点的值
	// 就是从根节点 root 为起点的最大单边路径和
	return max(leftMaxSum, rightMaxSum) + root.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
