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
