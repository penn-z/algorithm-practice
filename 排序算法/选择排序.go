/*
	选择排序:
		分已排序区间和未排序区间
		遍历未排序区间，每次从中选出最小元素，插入到已排序区间中
*/

package main

import "fmt"

func main() {
	arr := []int{7, 4, 1, 2, 8, 5, 11}
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		// 假设最小元素
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] { // 寻找最小的数
				min = j // 将最小数的索引保存
			}
		}

		// 放置到区间中
		arr[i], arr[min] = arr[min], arr[i]
	}

	return
}
