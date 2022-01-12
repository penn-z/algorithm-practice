package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
   层序遍历即可
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		levelNodeNum := len(queue)
		for i := 0; i < levelNodeNum; i++ {
			// 队首元素出队
			head := queue[0]
			queue = queue[1:]

			if head.Left != nil {
				queue = append(queue, head.Left)
			}

			if head.Right != nil {
				queue = append(queue, head.Right)
			}

			// 当前层次最右一个元素
			if i == levelNodeNum-1 {
				res = append(res, head.Val)
			}
		}
	}

	return res
}
