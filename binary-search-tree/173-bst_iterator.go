package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	arrs []int
}

func main() {

}

func Constructor(root *TreeNode) BSTIterator {
	bstIterator := new(BSTIterator)
	bstIterator.Inorder(root)
	return *bstIterator
}

func (this *BSTIterator) Inorder(node *TreeNode) {
	// base case
	if node == nil {
		return
	}

	// 中序遍历
	// 左
	this.Inorder(node.Left)

	// 中
	this.arrs = append(this.arrs, node.Val)

	// 右
	this.Inorder(node.Right)

	return
}

func (this *BSTIterator) Next() int {
	// 取出首个元素
	item := this.arrs[0]
	this.arrs = this.arrs[1:]
	return item
}

func (this *BSTIterator) HasNext() bool {
	return len(this.arrs) > 0
}

/*
	思路2:
		迭代版
		1. 中序遍历--迭代
		2. 此时无需预先计算出中序遍历的全部结果，只需要实时维护当前的栈情况即可
*/
type BSTIteratorV2 struct {
	stack   []*TreeNode
	curNode *TreeNode
}

func ConstructorV2(root *TreeNode) BSTIteratorV2 {
	return BSTIteratorV2{
		curNode: root,
	}
}

func (this *BSTIteratorV2) NextV2() int {
	// 这里注意，不再主动遍历stack，调用NextV2即可遍历

	// base case

	// this.curNode =
	for this.curNode != nil {
		// 入栈
		this.stack = append(this.stack, this.curNode)
		// 左
		this.curNode = this.curNode.Left
	}

	// 弹栈
	top := this.stack[len(this.stack)-1]
	val := top.Val // 中
	this.stack = this.stack[:len(this.stack)-1]
	this.curNode = top.Right // 右
	return val
}

func (this *BSTIteratorV2) HasNextV2() bool {
	return this.curNode != nil || len(this.stack) > 0
}
