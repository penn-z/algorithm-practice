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

/*
	快慢指针，快指针每次走两步，慢指针每次走一步。
	若链表中存在换，则快慢指针则一定会相遇，即fast == slow, 返回true
*/
func hasCycleV2(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
