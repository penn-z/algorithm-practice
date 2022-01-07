package main

import "fmt"

func main() {
	// testRangeOfIndex()
	// maxInt()
	testStrNum()
}

func testRangeOfIndex() {
	// test 函数
	vals := []int{1}

	vals = vals[1:] // 不会越界
	fmt.Printf("test:%v\n", vals)
}

func maxInt() {
	// 最大值
	maxInt := int(^uint(0) >> 1)
	// 最小值
	minInt := int(^maxInt)
	fmt.Println(maxInt)
	fmt.Println(minInt)
}

func testStrNum() {
	num1 := "111"
	fmt.Println(num1[0] - '0')
}
