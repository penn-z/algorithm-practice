/*
	leetcode 43: 字符串相乘
		给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。

	思路: 找规律相乘
	例如 num1 = "123", num = "45"
					1 2 3     i
				x	  4 5     j
				----------
					  1	5
                    1 0
                  0 5
				    1 2
				  0	8       => [i+j, i+j+1]
				0 4
			    ----------
                0 5 5 3 5

		res:   [0,0,0,0,0,0]
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	num1 := "123"
	num2 := "45"
	res := multiply(num1, num2)
	fmt.Println("===res: ", res)

}

func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// 结果数组
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		// 竖式乘式上方数值
		n1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			// 算出与res低位值相加结果
			sum := res[i+j+1] + int(n1)*int(n2)
			// 求res低位实际值
			res[i+j+1] = sum % 10
			// 求res高位实际值
			res[i+j] += sum / 10 // 此处 > 10没问题，后面迭代计算时res[i+j+1]会处理进位
		}
	}

	// 把res转为string
	buff := strings.Builder{}
	for i := range res {
		if i == 0 && res[i] == 0 {
			continue
		}

		buff.WriteString(fmt.Sprintf("%d", res[i]))
	}

	return buff.String()
}
