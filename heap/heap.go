// 最小堆
package main

import (
	"fmt"
	"math"
)

// 最小堆
type MinHeap struct {
	Element  []int // 堆数组，从下标1开始储存数据
	HeapSize int   // 用于记录数组的大小，堆可以存储的最大数据个数
	RealSize int   // 记录堆中的已储存元素个数
}

func main() {
	minHeap := NewMinHeap(10)
	minHeap.insert(54)
	fmt.Println(minHeap.toString())
	minHeap.insert(43)
	fmt.Println(minHeap.toString())
	minHeap.insert(16)
	fmt.Println(minHeap.toString())
	minHeap.insert(7)
	fmt.Println(minHeap.toString())
	minHeap.insert(10)
	fmt.Println(minHeap.toString())

	minHeap.remove()
	fmt.Println(minHeap.toString())
}

func NewMinHeap(heapSize int) *MinHeap {
	h := &MinHeap{
		// 第一个元素不记录堆元素，仅用于结束insert中的for循环
		Element:  []int{math.MinInt64},
		HeapSize: heapSize,
		RealSize: 0,
	}

	return h
}

// 取堆顶元素
// 复杂度O(1)
func (h *MinHeap) peek() int {
	return h.Element[1]
}

// 添加元素
// 复杂度O(logN)
func (h *MinHeap) insert(value int) {
	// 是否大于HeapSize
	if h.RealSize >= h.HeapSize {
		// 堆满了
		fmt.Println("add too many elements!")
		return
	}

	h.RealSize++

	// 将添加的元素加到数组中
	h.Element = append(h.Element, value)
	// h.Element[h.RealSize] = value

	index := h.RealSize // 当前即将要调整的当前节点元素下标
	parent := index / 2
	for h.Element[index] < h.Element[parent] && parent > 0 { // 自下往上堆化
		h.Element[index], h.Element[parent] = h.Element[parent], h.Element[index]
		index = parent      // 当前元素下标上移
		parent = parent / 2 // 父元素下标也上移
	}
}

// 删除堆顶元素
// 复杂度O(logN)
func (h *MinHeap) remove() int {
	// 当前堆元素个数为0，返回
	if h.RealSize < 1 {
		fmt.Println("don't have any element!")
		return 0
	}

	// 获取堆顶元素
	topEle := h.Element[1]

	// 将堆中最后一个元素赋值给堆顶元素
	h.Element[1] = h.Element[h.RealSize]

	// 置空原堆中最后一个位置，截断堆数组即可
	h.Element = h.Element[:h.RealSize]

	h.RealSize-- // 减少元素长度
	index := 1   // 当前元素下标

	// 自上而下进行下沉
	for {
		minPos := index // 当次比较中循环(当前元素与左右孩子比较)，最小元素下标
		// 被删除节点的左孩子
		leftIndex := index * 2
		// 被删除节点的右孩子
		rightIndex := index*2 + 1

		// 当前节点与左孩子节点、右孩子节点进行比较，若大于左右孩子，则取左右孩子中较小值进行交换

		// 左孩子节点
		if leftIndex <= h.RealSize && h.Element[minPos] > h.Element[leftIndex] {
			minPos = leftIndex
		}

		// 右孩子节点
		if rightIndex <= h.RealSize && h.Element[minPos] > h.Element[rightIndex] {
			minPos = rightIndex
		}

		if minPos == index {
			// 最小元素下标没有改变，可知已无需交换元素下沉
			break
		}

		// 交换元素值
		h.Element[index], h.Element[minPos] = h.Element[minPos], h.Element[index]

		// 该次交换结束，最小元素游标下移，为下次迭代交换作准备
		index = minPos
	}

	return topEle
}

// 堆元素大小
func (h *MinHeap) size() int {
	return h.RealSize
}

func (h *MinHeap) toString() string {
	return fmt.Sprintf("%v", h.Element[1:])
}
