/*
	1. 从数列中挑出一个元素，称为 "基准"（pivot）;

	2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；

	3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{7, 4, 1, 2, 8, 5, 11}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	// 如果区域内的元素 少于 2个，退出递归
	if left >= right {
		return
	}

	// 将数组分区，并获得中间值的下标
	middleIndex := partition(arr, left, right)
	// middleIndex := randomPartition(arr, left, right)
	// middleIndex := partitionV2(arr, left, right)

	// 对左区域快速排序
	quickSort(arr, left, middleIndex-1)

	// 对右区域快速排序
	quickSort(arr, middleIndex+1, right)

}

func partition(nums []int, start, end int) int {
	// 取第一个数为基数
	// pivot, pivotIndex := arr[left], left
	pivot := nums[start]

	for start < end {
		// 从右往左找到第一个小于pivot的元素
		for start < end && nums[end] >= pivot {
			end--
		}

		nums[start] = nums[end] // 放到start位置

		// 从左往右找到第一个大于pivot的元素
		for start < end && nums[start] <= pivot {
			start++
		}

		nums[end] = nums[start] // 放到end位置
	}

	//最后基准放到start的位置，此时nums[]start]空了
	nums[start] = pivot
	return start
}

// 选择排序的划分区间思想，划分为『比pivot小的元素区间』与『比pivot大的元素区间』
func partitionV2(arr []int, left, right int) int {
	// 基准pivot左边的元素比其小，右边的元素比其大
	// 选择最后一个元素为基准
	// 利用选怎排序的思想，将arr[left....right-1]分为『比pivot小元素区间』与『比pivot大元素区间』

	// 选择基准
	pivot := arr[right]

	i := left
	for j := left; j <= right-1; j++ {
		// 如果当前 『比pivot大元素区间』的j元素小于pivot，则需要交换i, j
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			// 『比pivot小元素区间』向右扩张1，刚好至『比pivot大的元素区间』
			i = i + 1
		}
	}

	// 最后需要把基准元素与i置换，因为此时i恰好在第一个比pivot大的元素索引位置，也即『比pivot大的元素区间』初始位置
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// 随机数优化
func randomPartition(nums []int, left int, right int) int {
	randomEle := rand.Intn(right-left+1) + left
	nums[left], nums[randomEle] = nums[randomEle], nums[left]
	return partition(nums, left, right)
}
