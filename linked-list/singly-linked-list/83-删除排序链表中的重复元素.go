package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	一个指针遍历，记录pre和上一个不重复值
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 上一个不重复节点值
	maxInt := int(^uint(0) >> 1)
	minInt := int(^maxInt)
	lastNotRepeatedVal := minInt

	cur := head
	var pre *ListNode
	for cur != nil {
		// 重复
		if lastNotRepeatedVal == cur.Val {
			// 删除当前节点
			pre.Next = cur.Next
		} else {
			// 不重复，更新pre，更新lastNotRepeatedVal
			lastNotRepeatedVal = cur.Val
			pre = cur
		}

		cur = cur.Next
	}

	return head
}

// 快慢指针，fast每次前进1步探路，如果与slow不相同(及元素不重复)，则slow前进至fast指针处
func deleteDuplicatesV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}

		fast = fast.Next
	}

	// slow断开与后面重复元素的连接
	slow.Next = nil
	return head
}
