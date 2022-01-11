package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	// 依照 左 -> 右，上 -> 下，右 -> 左，下 -> 上 ... 顺序重复遍历
	if len(matrix) == 0 {
		return res
	}

	m, n := len(matrix), len(matrix[0])
	// 定义上下左右初始值
	up, down, left, right := 0, m-1, 0, n-1
	for {
		// 左 -> 右
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}

		// up下移一行，准备下个方向遍历
		up++
		if up > down {
			// 越界
			break
		}

		// 上 -> 下
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}

		// right左移一列，准备下个方向遍历
		right--
		if right < left {
			// 越界
			break
		}

		// 右 -> 左
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}

		// down上移一行，准备下个方向遍历
		down--
		if down < up {
			// 越界
			break
		}

		// 下 -> 上
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}

		// left右移一行，准备下个方向遍历
		left++
		if left > right {
			// 越界
			break
		}
	}

	return res
}
