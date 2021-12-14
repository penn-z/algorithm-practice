package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 24345
	res := translateNum(num)
	fmt.Println("===res: ", res)
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}

	if num < 26 {
		return 2
	}

	str := strconv.Itoa(num)

	// dp[i]表示第i个数字时，可能翻译成字符串的方法数量

	// 递推公式
	/*
		dp[i]:
			若str[i-1]str[i]不可以拼接，则 dp[i] = dp[i-1]
			若str[i-1]str[i]可以贫瘠，则dp[i] = dp[i-2] + dp[i-1]
	*/

	// 初始化
	dp := make([]int, len(str))
	dp[0] = 1
	if str[:2] >= "10" && str[:2] <= "25" {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(str); i++ {
		// str[i-1]str[i]可以拼接
		if str[i-1:i+1] >= "10" && str[i-1:i+1] <= "25" {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(str)-1]
}
