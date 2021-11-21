package main

type MyQueue struct {
	InStack  []int
	OutStack []int
}

func main() {

}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.InStack = append(this.InStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.OutStack) == 0 {
		this.Tidy()
	}

	// 弹出outStack元素
	if len(this.OutStack) > 0 {
		outLen := len(this.OutStack)
		top := this.OutStack[outLen-1]
		this.OutStack = this.OutStack[0 : outLen-1]
		return top
	}

	return 0
}

// 获取队列开头元素
func (this *MyQueue) Peek() int {
	res := this.Pop()
	if res != 0 {
		this.OutStack = append(this.OutStack, res)
	}

	return res
}

func (this *MyQueue) Empty() bool {
	if len(this.OutStack) == 0 && len(this.InStack) == 0 {
		return true
	}

	return false
}

// 整理队列
func (this *MyQueue) Tidy() {
	for len(this.InStack) > 0 {
		this.OutStack = append(this.OutStack, this.InStack[len(this.InStack)-1])
		this.InStack = this.InStack[:len(this.InStack)-1]
	}
}
