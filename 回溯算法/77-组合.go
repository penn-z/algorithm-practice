package main

import "fmt"

func main() {
	n, k := 4, 2
	ret := combine(n, k)
	fmt.Println("ret: ", ret)
}

var res = [][]int{}

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n == 0 || k == 0 {
		return res
	}

	path := []int{}
	startIndex := 1 // 遍历开始的下标
	backtrack(n, k, path, startIndex)

	return res
}

// path已选择的路径
// startIndex 遍历开始的索引下标
func backtrack(n, k int, path []int, startIndex int) {
	// 结束条件
	if len(path) == k {
		// 加入路径
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
		return
	}

	// for 选择 in 选择列表:
	for i := startIndex; i <= n; i++ {
		// 排除不合法路径，该题无需该动作

		// 选择，加入路径
		path = append(path, i)

		// 递归
		backtrack(n, k, path, i+1)

		// 撤销选择，退出路径
		path = path[:len(path)-1]
	}

	return
}
