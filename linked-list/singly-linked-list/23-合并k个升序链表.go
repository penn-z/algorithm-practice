package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 顺序合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{
		Next: lists[0],
	}

	cur := dummy
	for i := 0; i < len(lists); i++ {
		cur.Next = mergeTwoListRecusion(cur.Next, lists[i])
	}

	return dummy.Next
}

// 分治合并
func mergeKListsV2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergePartition(lists, 0, len(lists)-1)
}

// 分区合并
func mergePartition(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	return mergeTwoListRecusion(mergePartition(lists, left, mid), mergePartition(lists, mid+1, right))
}

// 递归合并两个链表
func mergeTwoListRecusion(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListRecusion(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListRecusion(l2.Next, l1)
		return l2
	}
}
