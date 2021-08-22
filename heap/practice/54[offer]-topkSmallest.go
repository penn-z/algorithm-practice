/*
   思路:
       很明显是求TopK小元素
       1. 构建最小堆，依次将所有元素加入到堆中(使用官方包即可)
       2. 遍历从最小堆中弹出K原元素，加入到结果集T中，结果集T中的元素即是TopK小的元素
*/

package main

import (
	"container/heap"
	"fmt"
)

// 要使用官方container/heap包，需要实现sort, push, pop接口
type List []int

func main() {
	arr := []int{3, 2, 1}
	res := getLeastNumbers(arr, 2)
	fmt.Printf("res: %+v\n", res)
}

func getLeastNumbers(arr []int, k int) []int {
	list := List(arr)
	heap.Init(&list)
	res := []int{}
	for i := 1; i <= k; i++ {
		res = append(res, heap.Pop(&list).(int))
	}

	return res
}

func (l List) Len() int {
	return len(l)
}

// 最小堆
func (l List) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *List) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

func (l *List) Push(v interface{}) {
	*l = append(*l, v.(int))
}
