package main

import "fmt"

func main() {
	arr := []int{7, 4, 1, 2, 8, 5, 11}
	mergeSort(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) {
	mergeC(arr, 0, len(arr)-1)
}

func mergeC(arr []int, left, right int) {
	if left >= right {
		return
	}

	// 取p到r之间的中间位置q
	mid := left + (right-left)/2

	// 分治递归
	// 左半区域
	mergeC(arr, left, mid)
	// 右半区域
	mergeC(arr, mid+1, right)

	// 合并
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	// 申请一个与arr[left....right]大小一样的数组
	tmp := make([]int, right-left+1)
	// i为左半数组的第一个元素, j为右半数组的第一个元素
	i, j := left, mid+1
	// k为tmp的下标
	k := 0
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			// 左半数组下标右移
			i++
		} else {
			tmp[k] = arr[j]
			// 右半数组下标右移
			j++
		}

		k++
	}

	// 判断哪个数组还有剩余元素
	start, end := i, mid // 先假设前半数组还有剩余元素
	if j <= right {
		start = j
		end = right
	}

	// 剩余数据拷贝到tmp
	// 这里需要取 =，因为上面的for循环遍历中，最后一次可能是i++或j++，可能出现 i == mid+1, 或者j == right+1情况
	for start <= end {
		tmp[k] = arr[start]
		start++
		k++
	}

	// tmp数组拷贝到arr
	for i := 0; i <= right-left; i++ {
		// 这里需要left+i，是因为
		// 直接操作arr底层数组内存
		arr[left+i] = tmp[i]
	}

	return
}
