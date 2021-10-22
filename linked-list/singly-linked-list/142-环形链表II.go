/*
	leetcode 142: 环形链表II
	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
	说明：不允许修改给定的链表。

	思路1:
		遍历链表，使用map记录每个结点指针地址，若重复出现，则返回该节点
		空间复杂度: O(n)

	思路2:
		1. 定义快慢指针fast, slow，快指针每次前进2步，慢指针前进1步。
		2. 若存在环，则快慢指针一定在环中相遇，且fast一定比slow在环中多走了n(n>=1)圈，且fast走过的路程是slow 2倍。
		3. 假设从链表头到环入口距离为x, 两指针在环中相遇时，slow在环中走过距离为y, y至回到环入口距离为z
			则有 slow = x + y
				fast = x + n(y + z) + y，其中n为fast在环中走过的圈数
				又因为 fast = 2 * slow
				化简得 x = n(y+z) - y
				右边提取一个(y+z)，则x = (n-1)(y+z) + z
			可知，当快慢指针在环中相遇时，可定义两个新指针index1, index2，index1从相遇点触发，index2出链表头触发，最终相遇时即为环入口。此时index1在环中走了(n-1)圈
		参考:
			https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.md
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	// 定义map，记录结点指针地址
	pointerMap := make(map[*ListNode]bool)

	// 虚拟节点
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy
	for cur.Next != nil {
		if _, ok := pointerMap[cur.Next]; ok {
			return cur.Next
		}

		pointerMap[cur.Next] = true
		cur = cur.Next
	}

	return nil
}

func detectCycleV2(head *ListNode) *ListNode {
	// 定义快慢指针
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 快慢指针相遇时
		if fast == slow {
			index1, index2 := slow, head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}

			return index1
		}
	}

	return nil
}
