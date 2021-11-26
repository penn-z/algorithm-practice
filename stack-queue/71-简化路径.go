/*
	leetcode 71: 简化路径

	思路: 利用栈特性
	1. 用""/"分隔符把path分隔为数组
	2. 遍历数组，遇到 ".", "..", ""时需要分支处理，其余字符串压入栈中
		".": 表示当前路径，不做任何事，跳过
		"..": 表示返回上一级，需要把栈顶元素弹出
		"": 无需做任何事
	3. 遍历结束后，遍历栈元素，每个元素之间使用/拼接即可
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// path := "/../"
	path := "/home/"
	res := simplifyPath(path)
	fmt.Println("res: ", res)
}

func simplifyPath(path string) string {
	res := ""
	if len(path) == 0 {
		return res
	}

	dirStack := []string{}

	strArr := strings.Split(path, "/")
	fmt.Println("strArr: ", strArr)
	for _, str := range strArr {
		switch str {
		case ".", "":
			// do nothing
		case "..":
			// 栈顶元素弹出
			if len(dirStack) > 0 {
				dirStack = dirStack[:len(dirStack)-1]
			}
		default:
			// 压入栈中
			dirStack = append(dirStack, str)
		}
	}

	res = strings.Join(dirStack, "/")
	return "/" + res
}
