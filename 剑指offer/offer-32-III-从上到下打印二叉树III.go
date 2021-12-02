package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

/*
   思路: BFS-层序遍历
       1. 依然是层序遍历，利用队列进行BFS
       2. 外循环控制层，内循环遍历每层的兄弟节点
       3. 每层遍历完成后，如果当前层是偶数层，则当前层的数组结果需要反转后再加入到最终结果二维数组中
       4. 输出结果数组即可
*/
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{}
	// 根节点入队
	queue = append(queue, root)
	for len(queue) > 0 {
		levelRes := []int{}
		levelLen := len(queue)
		for i := 0; i < levelLen; i++ {
			// 队头元素出队
			tail := queue[0]
			queue = queue[1:]

			if tail.Left != nil {
				// 左子树入队
				queue = append(queue, tail.Left)
			}

			if tail.Right != nil {
				// 右子树入队
				queue = append(queue, tail.Right)
			}

			levelRes = append(levelRes, tail.Val)
		}

		// 当前层为偶数层, 需要反转数组
		// len(res) + 1是因为要算上该层
		if (len(res)+1)%2 == 0 {
			levelRes = reverseArr(levelRes)
		}

		res = append(res, levelRes)
	}

	return res
}

// 反转数组
func reverseArr(arr []int) []int {
	// 头尾双指针反转数组
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
