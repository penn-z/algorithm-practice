package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	// 定义map，记录结点指针地址
	pointerMap := make(map[*ListNode]bool)

	// 虚拟节点
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy
	for cur.Next != nil {
		if _, ok := pointerMap[cur.Next]; ok {
			return true
		}

		pointerMap[cur.Next] = true
		cur = cur.Next
	}

	return false
}
