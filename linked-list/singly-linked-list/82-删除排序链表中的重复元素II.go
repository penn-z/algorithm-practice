/*
	存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
	返回同样按升序排列的结果链表。

	输入：head = [1,2,3,3,4,4,5]
	输出：[1,2,5]

	输入：head = [1,1,1,2,3]
	输出：[2,3]


	一次遍历，定义前置节点pre, 当前节点cur
        比较cur.Val = cur.Next.Val，如果相等，则cur继续前移，直到不相等为止，此时cur位于最后一个重复元素。
        判断pre.Next是否与cur相等
            若相等，说明cur没有进行内循环移动，无重复元素，则pre向pre.Next移动即可
            若不相等，pre.Next指向cur.Next，pre指针不移动
        cur移动到cur.Next，随后继续外层大循环
*/
package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head
	for cur != nil {
		// 跳过重复节点，使得当前cur指向最后一个重复元素
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next == cur {
			// cur没有进入内循环移动
			// pre和cur之间没有重复节点，pre后移
			pre = pre.Next
		} else {
			//pre->next指向cur的下一个位置（相当于跳过了当前的重复元素）
			//但是pre不移动，仍然指向已经遍历的链表结尾
			pre.Next = cur.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}
