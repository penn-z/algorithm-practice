/**
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

输入：head = [1,2,3,4]
输出：[1,4,2,3]

输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

 快慢指针
        1. 快慢指针从头结点开始遍历，慢指针每次前进1， 快指针前进2
        2. 当快指针到尾结点时，慢指针在中点，往后慢指针继续前进，开始翻转后半链表
        3. 最后前半链表和后半链表错位相交，前半链表结点先，后边链表后，交错连接

*/

package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	// 取链表中间左边结点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 此时slow已到中点，后半链表需要翻转
	mid := slow.Next
	// 这里需要把原链表后半链表置为nil
	slow.Next = nil

	cur := mid
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	// pre为已翻转好的后半链表

	// 目前就两链表pre, head
	cur = head
	for pre != nil && cur != nil {
		tmp1 := cur.Next
		tmp2 := pre.Next

		cur.Next = pre
		cur = tmp1

		pre.Next = cur
		pre = tmp2
	}

	return
}

/**
slice记录链表，
然后定义快慢指针i, j，按规律从slice中取出node连接
*/
func reorderListV2(head *ListNode) {
	if head == nil {
		return
	}

	// 用slice存储链表
	nodeSlice := []*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodeSlice = append(nodeSlice, cur)
	}

	// 双指针，按规律连接链表
	i, j := 0, len(nodeSlice)-1
	for i < j {
		nodeSlice[i].Next = nodeSlice[j]
		i++
		if i == j {
			break
		}

		nodeSlice[j].Next = nodeSlice[i]
		j--
	}

	// 需要把中间结点Next置0，否则会循环
	nodeSlice[i].Next = nil

	return
}
