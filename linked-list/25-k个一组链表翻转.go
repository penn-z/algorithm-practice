package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	// 区间[a, b)包含k个元素
	a, b := head, head
	for i := 0; i < k; i++ {
		// 如果不足k个元素，则不需要翻转
		if b == nil {
			return head
		}

		b = b.Next
	}

	// 翻转前k个元素
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 翻转[a, b)之间的链表元素
func reverse(a, b *ListNode) *ListNode {
	cur := a
	// next := a
	var pre *ListNode
	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
