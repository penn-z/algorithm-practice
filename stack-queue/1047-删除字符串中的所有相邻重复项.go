/*
	leetcode 1047:
		给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
		在 S 上反复执行重复项删除操作，直到无法继续删除。
		在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

	demo1:
		input: "abbaca"
		outout: "ca"

	思路: 使用栈LIFO特性
		1. 定义一个栈，用于储存相邻不重复字符
		2. 遍历字符串，与栈顶字符不相同则入栈，相同则删除当前字符及栈顶字符
		3. 遍历结束后，把栈(字节数组)转为字符串即可

*/
package main

func main() {

}

/*
   思路: 使用栈LIFO特性
   1. 定义一个栈，用于储存相邻不重复字符
   2. 遍历字符串，与栈顶字符不相同则入栈，相同则删除当前字符及栈顶字符
   3. 遍历结束后，把栈(字节数组)转为字符串即可
*/
func removeDuplicates(s string) string {
	stack := []byte{}
	for i, _ := range s {
		// 当前空栈，直接入栈即可
		if len(stack) < 1 {
			stack = append(stack, s[i])
			continue
		}

		top := stack[len(stack)-1]
		if s[i] == top {
			// 栈顶元素出栈
			stack = stack[:len(stack)-1]
		} else {
			// 字符与栈顶元素不相同，则入栈
			stack = append(stack, s[i])
		}
	}

	// 遍历结束
	return string(stack)
}
