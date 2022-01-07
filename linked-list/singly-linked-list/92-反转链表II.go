package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代法，哨兵节点，部分翻转
func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil {
		return nil
	}

	if left == right {
		return head
	}

	// 定义哨兵节点
	dummy := &ListNode{
		Next: head,
	}

	// 找到left节点前一节点
	preLeft := dummy
	for i := 1; i < left; i++ {
		preLeft = preLeft.Next
	}

	// 找到rightNode节点
	rightNode := preLeft
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 记录leftNode
	leftNode := preLeft.Next
	// 记录rightNode.Next
	rightNodeNext := rightNode.Next

	// 切断left到right的链表
	preLeft.Next = nil   // 先删最前面 preLeft.Next -> leftNode的引用
	rightNode.Next = nil // 再删最后面rightNode -> rightNode.Next的引用

	// 翻转[leftNode, rightNode]链表
	reverse(leftNode, rightNode.Next)

	// 重新拼接
	preLeft.Next = rightNode
	leftNode.Next = rightNodeNext

	return dummy.Next
}

// 翻转区间[a, b)之间的链表元素
func reverse(a, b *ListNode) *ListNode {
	cur := a
	var pre *ListNode
	for cur != b {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
