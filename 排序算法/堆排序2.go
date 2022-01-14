package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 2, 4, 7, 1, 5, 3}
	arr = sort(arr)
	fmt.Println(arr)
}

// 建堆
// 大顶堆，自上而下堆化
// 当前元素为i，左孩子为2i, 右孩子为2i+1
func buildHeap(arr []int, arrLen int) {
	// 从最后一个非叶子节点开始，依次堆化
	for i := arrLen/2 - 1; i >= 0; i-- {
		heapify(arr, arrLen, i)
	}

}

// 堆化
func heapify(arr []int, arrLen int, i int) {
	for {
		// 假设当前元素与左右孩子节点中，当前元素最大，maxPos为当前最大元素索引
		maxPos := i
		leftChild := 2 * i
		rightChild := 2*i + 1

		// maxPos与左孩子比较大小
		if leftChild < arrLen && arr[maxPos] < arr[leftChild] {
			maxPos = leftChild
		}

		// maxPos与右孩子比较大小
		if rightChild < arrLen && arr[maxPos] < arr[rightChild] {
			maxPos = rightChild
		}

		// maxPos位置无变化，说明i堆化完成
		if maxPos == i {
			break
		}

		// 置换元素值
		swap(arr, i, maxPos)

		// 更新i位置, i往下移
		i = maxPos
	}
}

// 排序
func sort(arr []int) []int {
	arrLen := len(arr)
	// 建堆
	buildHeap(arr, arrLen)
	// 排序，不断把堆顶元素弹出，并保持堆化的过程
	for i := arrLen - 1; i >= 0; i-- {
		/*
			堆顶元素与最后一个元素交换位置
				如果是大顶堆，则最后的数组从右往左降序，则正序为升序
				小顶堆反之，正序为降序数组
		*/
		swap(arr, 0, i)

		// 数组长度减小
		arrLen--

		// 保持数组堆化 (堆顶元素是最后元素置换上来的，故需要对堆顶元素堆化操作)
		heapify(arr, arrLen, 0)
	}

	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
