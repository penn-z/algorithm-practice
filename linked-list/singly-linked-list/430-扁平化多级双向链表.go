/*
	leetcode 430: 扁平化多级双向链表

	多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
	给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。

	demo1:
		输入: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
		输出: [1,2,3,7,8,11,12,9,10,4,5,6]


	思路1:
		递归
		1. 当前节点 headhead 没有 childchild 节点：直接让指针后即可，即 head = head.nexthead=head.next；
		2. 当前节点 headhead 有 childchild 节点：将 head.childhead.child 传入 flatten 函数递归处理，拿到普遍化后的头结点 cheadchead，然后将 headhead 和 cheadchead 建立“相邻”关系（注意要先存起来原本的 tmp = head.nexttmp=head.next 以及将 head.childhead.child 置空），然后继续往后处理，直到扁平化的 cheadchead 链表的尾部，将其与 tmptmp 建立“相邻”关系。
*/

package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func main() {

}

// 递归: 详见说明
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	dummy := &Node{}
	dummy.Next = root
	cur := dummy.Next
	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
		} else {
			tmp := cur.Next
			// 获取child节点组成的链表头节点childHead
			childHead := flatten(cur.Child)
			// 获取child节点组成的链表的尾结点childEnd
			childEnd := getEnd(childHead)

			// child链表尾结点next指向cur.Next
			childEnd.Next = cur.Next
			// childHead上一节点指向当前节点cur
			childHead.Prev = cur
			if cur.Next != nil {
				cur.Next.Prev = childEnd
			}
			// 当前节点指向childHead
			cur.Next = childHead
			// 置空cur.Child
			cur.Child = nil

			// cur前进至原本的cur.Next
			cur = tmp
		}
	}

	return dummy.Next
}

// 获取尾部节点
func getEnd(node *Node) *Node {
	for node.Next != nil {
		node = node.Next
	}

	return node
}

// func flattenIteration(root *Node) *Node {

// }
