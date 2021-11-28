package main

/*
   思路: 辅助栈
   1. 定义一个普通栈outStack用于pop()操作，一个辅助最小栈minStack用于min()操作
   2. push()时，除了往outStack入栈，还需要判断值是否小于minStack栈顶元素值，是则往minStack也入栈
   3. pop()时，除了outStack栈顶元素出栈，还需要判断outStack栈顶元素是否等于minStack栈顶元素，是则minStack栈顶元素出栈
   4. top()为outStack元素值, min()为minStack栈顶元素值
*/

type MinStack struct {
	outStack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.outStack = append(this.outStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		if val <= this.minStack[len(this.minStack)-1] {
			this.minStack = append(this.minStack, val)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.outStack) > 0 {
		top := this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
		if top == this.minStack[len(this.minStack)-1] {
			this.minStack = this.minStack[:len(this.minStack)-1]
		}
	}
}

func (this *MinStack) Top() int {
	return this.outStack[len(this.outStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
