/*
	给定一个只包括'(', ')', '{', '}', '[', ']'的字符串s，判断字符串是否有效。
	有效字符串需满足:
		1. 左括号必须用相同类型的有括号闭合
		2. 左括号必须以正确的顺序闭合

	demo1:
		input: s = "()"
		output: true

	demo2:
		input: s = "()[]{}"
		output: true

	demo3:
		input: s = "(]"
		output: false

	demo4:
		input: s = "{[]}"
		output: true

	思路:
		1. 使用栈LIFO特性，维护一个栈
		2. 遍历字符串，
			1. 遇到左括号就入栈
			2. 遇到右括号
				若栈为空，说明是独立的右括号，不合法
				若栈不为空， 则取出栈顶元素匹配，匹配则继续，否则返回false退出
		3. 遍历结束，若栈不为空，说明有独立的左括号，返回false；否则全匹配，返回true
*/
package main

func main() {

}

var (
	bracketsMap = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
)

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	bracketsStack := []string{}
	for _, char := range s {
		_, ok := bracketsMap[string(char)]
		if ok {
			// 左括号，入栈
			bracketsStack = append(bracketsStack, string(char))
		} else {
			// 右括号，取出栈顶元素进行匹配
			rightCur := string(char)

			// 此时栈为空，说明先入了个右括号，不合法
			if len(bracketsStack) == 0 {
				return false
			}

			// 取出栈顶左括号
			top := bracketsStack[len(bracketsStack)-1]
			right := bracketsMap[top]
			if right != rightCur {
				return false
			}

			bracketsStack = bracketsStack[:len(bracketsStack)-1]
		}
	}

	if len(bracketsStack) > 0 {
		return false
	}

	return true
}
