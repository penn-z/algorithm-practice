/*
	leetcode 234: 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

	思路1: 使用栈特性+快慢指针
		1. 快慢指针找右中点，慢指针走过的节点值依次入栈
		2. 找到右中点后，若fast指针指向nil，则原链表节点数为偶数，此时栈顶元素为后半链表第一个节点的值，需要出栈剔除。
		3. 慢指针继续前进，当前值与栈顶元素比较。


	思路2: 使用快慢指针+后半链表翻转比较
		1. 快慢指针找到后半链表，反转后，再用前后指针向中间移动，比较指针所在节点的值

*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func isPalindrome(head *ListNode) bool {
	// 快慢指针
	fast, slow := head, head
	// 维护一个栈，记录前半部分链表的节点值
	valStack := []int{slow.Val}
	// 中点有两个时，说明链表节点数为偶数，此时取右中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		valStack = append(valStack, slow.Val)
	}

	if fast == nil {
		// 快指针指向链表尾部nil指针，则链表节点数为偶数，次数删除栈顶一元素即可
		valStack = valStack[:len(valStack)-1]
	}

	// 慢指针已到中点
	for slow != nil {
		// 栈顶元素弹出
		top := valStack[len(valStack)-1]
		valStack = valStack[:len(valStack)-1]

		if top != slow.Val {
			return false
		}

		slow = slow.Next
	}

	return true
}

/*
	快慢指针+反转后半链表
	1. 翻转链表的后一半，然后双指针，一个指针指向前半部分的头部，另一个指向翻转后的一半的结尾，判断元素是否相等。
	2. 使用快慢指针，获取到中间结点，从中间结点开始翻转后一半链表。
	3. 再使用前后双指针，向链表中间移动，一次比较两指针所在节点的值，不相等则返回false退出。
*/
func isPalindromeV2(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到链表中点
	halfNode := findHalfNode(head)

	// 后半链表翻转，得到翻转后的head
	endHead := reverseList(halfNode.Next)

	// 此时定义左右边界指针，向中间靠近
	p1 := head
	p2 := endHead

	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}

		// 移动左右指针
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

// 找到中间结点
func findHalfNode(head *ListNode) *ListNode {
	// 快慢指针
	fast, slow := head, head
	// ps: 这里要注意，中间有两个时，取左边那个，因为后面翻转后半链表需要传入node.Next，故而这里取左中点
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// 翻转链表
func reverseList(head *ListNode) *ListNode {
	// 定义pre先驱结点，cur当前结点
	cur := head
	var pre = &ListNode{}
	for cur != nil {
		temp := cur.Next // 暂存cur下一结点，以便cur移动
		cur.Next = pre

		// 移动pre, cur
		pre = cur
		cur = temp
	}

	return pre
}
