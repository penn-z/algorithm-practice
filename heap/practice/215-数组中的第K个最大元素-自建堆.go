/**
  思路:
  求第k个最大的元素，可以用小顶堆来实现
  1. 把数组通过堆化建立长度大小为k的小顶堆
  2. 依次将元素添加到小顶堆中
  3. 当堆大小 == k时，将当前元素与堆顶元素比较:
      1. 若当前元素 <= 堆顶元素，则放弃当前元素，继续循环
      2. 若当前元素 > 堆顶元素，则删除堆顶元素，将当前元素加入到堆中
  4. 最后小顶堆的堆顶元素即为第k个最大的元素
*/

package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	res := findKthLargest(nums, k)
	fmt.Println("===res: ", res)
}

func findKthLargest(nums []int, k int) int {
	// 前k个元素用来建堆
	heap := nums[:k]
	builHeap(heap, k)

	// builHeap(nums, len(nums))

	fmt.Println("heap:", heap)
	// return 0
	for i := k; i < len(nums); i++ {
		fmt.Println("i: ", i)
		if nums[i] > peek(heap) {
			// 加入堆中
			add(heap, nums[i], k)
		}
	}

	return peek(heap)
}

// 建堆，小顶堆
func builHeap(nums []int, arrLen int) {
	// 自上而下堆化，从最后一个非叶子节点开始
	for i := arrLen / 2; i >= 0; i-- {
		heapify(nums, arrLen, i)
	}
}

// 堆化
func heapify(nums []int, arrLen, i int) {
	for {
		// 最小值索引
		minPos := i
		// 左右孩子
		leftChild, rightChild := 2*i, 2*i+1
		if leftChild < arrLen && nums[i] > nums[leftChild] {
			minPos = leftChild
		}

		if rightChild < arrLen && nums[minPos] > nums[rightChild] {
			minPos = rightChild
		}

		// minPos未改变，则i堆化完成
		if minPos == i {
			break
		}

		// 交换i, minPos位置
		swap(nums, i, minPos)

		// i指针下移
		i = minPos
	}
}

func add(nums []int, newVal, arrLen int) {
	// 删除堆顶元素
	nums[0] = newVal
	// 保持堆化
	heapify(nums, arrLen, 0)
}

func peek(nums []int) int {
	return nums[0]
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
