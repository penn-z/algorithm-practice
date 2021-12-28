/*
	插入排序:
	核心思想:
		取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。
		重复这个过程，直到未排序区间中的元素为空

		1. 将数组分为两个区间: 已排序区间和未排序区间，初始的已排序区间只有1个元素，即数组第一个元素
		2. 遍历未排序数组，每次从下标指向的元素取出元素，插入到已排序区间中，期间可能需要挪动已排序区间元素的位置

*/
package main

import "fmt"

func main() {
	arr := []int{7, 4, 1, 2, 8, 5, 11}
	inserSort(arr)
	fmt.Println(arr)
}

func inserSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 从第2个元素开始遍历，即未排序区间开始
	for i := 1; i < len(arr); i++ {
		curValue := arr[i] // 当前要被插入到已排序区间的元素
		// 已排序区间的末端位置
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if arr[j] > curValue {
				// 当前a[j]数据往后移一位
				arr[j+1] = arr[j] // 数据移动
			} else {
				break
			}
		}

		// 插入到已排序区间后一位，for循环中最后会执行j--，导致curValue实际要待的位置向左了一位
		arr[j+1] = curValue
	}
}
