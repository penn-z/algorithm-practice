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

	/*
		优化重点:
			因为每层我们都是从startIndex起始位置开始遍历的，但是在 (startIndex+1)之后的数都不足以 满足我们需要的元素个数了，
			可以不必再搜索，进行剪枝操作

		优化过程:
			1. 已选择的元素个数: len(path)
			2. 还需要选择的元素个数: k - len(path)
			3. 在集合n中至多要从该起始位置: n - (k - len(path)) + 1 开始搜索，才满足题意条件
	*/

	// for 选择 in 选择列表:
	// i为本次搜索的起始位置
	// for i := startIndex; i <= n; i++ { // 剪枝前
	for i := startIndex; i <= n-(k-len(path))+1; i++ { // 剪枝后
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
