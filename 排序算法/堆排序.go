package main

import "fmt"

func main() {
	arr := []int{8, 2, 4, 7, 1, 5, 3}
	arr = sort(arr)
	fmt.Println(arr)
}

// 大顶堆，自上而下堆化
// 此处的arr是从索引0开始的，索引0不是占位符
// 则当前元素为i，左孩子为2i, 右孩子为2i+1
func buildHeap(arr []int, arrLen int) {
	// 从最后一个非叶子节点开始，一次堆化
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, arrLen, i)
	}

	fmt.Println("建堆完成，arr: ", arr)
}

// 堆化
func heapify(arr []int, arrLen int, i int) {
	for {
		// 假设当前元素与左右孩子节点中，当前元素最大，则maxPos是当前元素的索引
		maxPos := i
		leftChild := 2 * i
		rightChild := 2*i + 1
		// 如果当前元素比其左孩子小，则maxPost需要替换为左孩子
		if leftChild < arrLen && arr[i] < arr[leftChild] {
			maxPos = leftChild
		}

		// rightChild > max(i, leftchild)，则maxPos取rightChild
		if rightChild < arrLen && arr[maxPos] < arr[rightChild] {
			maxPos = rightChild
		}

		// maxPos没有发生变化，说明当前i已经是最大元素，i堆化完成
		if maxPos == i {
			break
		}

		// 交换位置，即堆化
		swap(arr, i, maxPos)

		// 更新i位置，i往下移
		i = maxPos
	}
}

func sort(arr []int) []int {
	arrLen := len(arr)
	// 先堆化
	buildHeap(arr, len(arr))
	// 排序，不断把堆顶元素弹出，并保持堆化的过程
	for i := arrLen - 1; i >= 0; i-- {
		/*
			堆顶第一个元素与最后一个元素交换位置
				如果是大顶堆，则最后的数组从右往左降序，则正序为升序数组
				小顶堆反之，正序为降序数组
		*/
		swap(arr, 0, i)
		// 数组长度减小
		arrLen--
		// 保持堆化
		heapify(arr, arrLen, 0)
	}

	return arr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
