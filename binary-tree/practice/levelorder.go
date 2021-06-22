package main

import (
	"fmt"
)

// 层序遍历
// 广度优先，利用队列先进先出特点

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

	vals := levelorder(root)

	fmt.Printf("traversal vals:%v\n", vals)
}

// 层序遍历
func levelorder(root *TreeNode) (vals [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{}

	// 根节点入队
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}              // 记录当前层级队列
		currLevelEleLen := len(queue) // 当前层级元素长度

		for i := 0; i < currLevelEleLen; i++ {
			// 队头元素出队
			tmp := queue[0]
			queue = queue[1:]

			level = append(level, tmp.Val)
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}

			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}

		vals = append(vals, level)
	}

	return
}
