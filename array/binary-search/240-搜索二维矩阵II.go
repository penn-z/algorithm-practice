/*
	leetcode 240: 搜索二维矩阵II
	编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
	每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。

	demo1:
		input: [
			[1, 4, 7, 11, 15],
			[2, 5, 8, 12, 19],
			[3, 6, 9, 16, 22],
			[10, 13, 14, 17, 24],
			[18, 21, 23, 26, 30],
		]

		target = 5

		output: true

	思路1: 二分查找
		1. 遍历外层
		2. 遍历内层，二分查找


	思路2: BST搜索
		1. 从右上作为起点，向左方、下方开始搜索，可发现是一颗BST树
		2. 利用BST特点，当前节点比左孩子节点大，比右孩子节点小，往左下搜索
*/
package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	// 二分查找
	lenN := len(matrix)

	for i := 0; i <= lenN-1; i++ {
		arr := matrix[i]
		if len(arr) == 0 {
			break
		}

		left, right := 0, len(arr)-1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] == target {
				return true
			} else if arr[mid] < target {
				// 左边界右移
				left = mid + 1
			} else if arr[mid] > target {
				// 右边界左移
				right = mid - 1
			}
		}
	}

	return false
}

func searchMatrixV2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	left, right := n-1, 0
	for left >= 0 && right < m {
		if matrix[right][left] < target {
			// 指针向下移
			right++
		} else if matrix[right][left] > target {
			// 指针向左移
			left--
		} else {
			return true
		}
	}

	return false
}
