package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	res := largestValues(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	maxInt := int(^uint(0) >> 1)
	minInt := int(^maxInt)

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		curLevelMaxInt := minInt
		// 当前层节点数量
		curNodeNums := len(queue)
		for i := 0; i < curNodeNums; i++ {
			// 队头节点出队
			head := queue[0]
			queue = queue[1:]
			curLevelMaxInt = max(curLevelMaxInt, head.Val)

			// 左右子节点入队
			if head.Left != nil {
				queue = append(queue, head.Left)
			}

			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}

		res = append(res, curLevelMaxInt)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
