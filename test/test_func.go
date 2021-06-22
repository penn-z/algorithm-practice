package main

import "fmt"

func main() {
	// test 函数
	vals := []int{1}

	vals = vals[1:] // 不会越界
	fmt.Printf("test:%v", vals)
}
