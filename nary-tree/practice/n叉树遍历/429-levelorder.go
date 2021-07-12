package main

import "fmt"

/*
	leetcode: 429-n叉树的层序遍历

*/

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	res := levelOrder(root)
	fmt.Printf("=====res:%+v", res)
}

/*
   思路:
   1. 每层实例一个记录当前层的变量
   2. 初始化变量记录当前层节点数量
   2. BFS广度优先遍历即可
*/
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*Node{}
	// 当前节点数量
	curLevelNodeNums := 1

	// 根节点入栈
	queue = append(queue, root)
	for len(queue) > 0 {
		// 记录当前层元素slice
		levelItems := []int{}
		// 下一层节点数量
		nextLevelNodeNums := 0
		for i := 0; i < curLevelNodeNums; i++ {
			// 队头元素出队
			head := queue[0]
			queue = queue[1:]
			levelItems = append(levelItems, head.Val)
			nextLevelNodeNums += len(head.Children)
			for _, childNode := range head.Children {
				queue = append(queue, childNode)
			}
		}

		curLevelNodeNums = nextLevelNodeNums
		res = append(res, levelItems)
	}

	return res
}
