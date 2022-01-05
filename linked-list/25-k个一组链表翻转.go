package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
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

// 迭代
/*
	用栈，我们把 k 个数压入栈中，然后弹出来的顺序就是翻转的！
		这里要注意几个问题：
		第一，剩下的链表个数够不够 k 个（因为不够 k 个不用翻转）；
		第二，已经翻转的部分要与剩下链表连接起来。
*/
func reverseKGroupV2(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	stack := []*ListNode{}
	dummy := &ListNode{}
	p := dummy
	for {
		count := 0
		tmp := head
		for tmp != nil && count < k {
			// 入栈
			stack = append(stack, tmp)
			// 指针前进
			tmp = tmp.Next
			count++
		}

		// 最后一段区间，不足k个元素
		if count != k {
			p.Next = head
			// 退出即可
			break
		}

		// 栈不为空，出栈
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p.Next = top
			// p指针前进
			p = p.Next
		}

		// p.Next指向下一个区间
		p.Next = tmp
		// head也指向下一个区间
		head = tmp
	}

	return dummy.Next
}
