package main

import "fmt"

func main() {
	arr := []int{7, 4, 1, 2, 8, 5}
	// arr = bubbleSort(arr)
	arr = bubbleSortV2(arr)
	fmt.Println(arr)
}

/*
	冒泡排序:
	升序
		相邻两两元素比较，较大元素放置后面，遍历完一轮，则最大元素已在数组末端。
		下一轮轮从 len(arr) - i -1继续重复操作
*/
func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

/*
	优化:
		如果某一轮排序时，没有发生交换，说明数组已经有序，无需继续排序操作
		定义swapped变量记录是否有序
*/
func bubbleSortV2(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var swapped bool = true
	for i := 0; i < len(arr)-1; i++ {
		// 如果上一轮没有发生元素交换，说明已有序
		if !swapped {
			break
		}

		// 假设有序
		swapped = false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 发生元素交换，未到有序状态
				swapped = true
			}
		}
	}

	return arr
}
