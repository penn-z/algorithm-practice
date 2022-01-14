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
