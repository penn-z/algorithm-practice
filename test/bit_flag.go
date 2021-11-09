package main

import "fmt"

func main() {
	var num int32 = 27
	// fmt.Println("after num:", num>>1)
	// count := 0
	// res := []int32{}
	// i := 0
	// for num > 0 {
	// 	fmt.Println("num & 1: ", num&1)
	// 	if num&1 == 1 {
	// 		fmt.Println("当前位的值: ", 1<<i)
	// 		res = append(res, 1<<i)
	// 	}

	// 	count = count + (num & 1)
	// 	num >>= 1
	// 	fmt.Println("num: ", num)
	// 	i++
	// }

	// fmt.Println("整型二进制位数量: ", count)
	// fmt.Println("res: ", res)
	res := TransferAttrFlagBitToSlice(num)
	fmt.Println("res: ", res)
}

func TransferAttrFlagBitToSlice(attrFlag int32) (res []int32) {
	if attrFlag <= 0 {
		return
	}

	i := 0
	for attrFlag > 0 {
		// 当前位为1
		if attrFlag&1 == 1 {
			res = append(res, 1<<i)
		}

		// attrFlag右移一位
		attrFlag >>= 1
		i++
	}

	return
}
