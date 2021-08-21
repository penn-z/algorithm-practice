package main

import (
	"container/heap"
	"fmt"
)

type Test struct {
	Num int
}

type TestList []*Test

func main() {
	list := &TestList{
		{
			Num: 6,
		},
		{
			Num: 10,
		},
		{
			Num: 3,
		},
		{
			Num: 7,
		},
	}
	// 需要实现sort, pop, push三个interface
	heap.Init(list)
	heapStr := ""
	i := 0
	for {
		if i >= list.Len() {
			break
		}

		heapStr = fmt.Sprintf("%s#%d", heapStr, (*list)[i].Num)
		i++
	}

	fmt.Printf("heap:%s\n", heapStr)
}

// Len, Less, Swap为sort包三个接口
// sort.Len
func (t *TestList) Len() int {
	return len(*t)
}

// sort.Less
func (t *TestList) Less(i, j int) bool {
	return (*t)[i].Num <= (*t)[j].Num
}

// sort.Swap
func (t *TestList) Swap(i, j int) {
	(*t)[i].Num, (*t)[j].Num = (*t)[j].Num, (*t)[i].Num
}

// heap.Push
func (t *TestList) Push(val interface{}) {
	*t = append(*t, val.(*Test))
}

// heap.Pop
func (t *TestList) Pop() interface{} {
	old := *t
	*t = old[:len(old)-1]
	// 弹出最后一个元素
	return old[len(old)-1]
}
