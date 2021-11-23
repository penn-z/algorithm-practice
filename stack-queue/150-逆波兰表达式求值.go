/*
	leetcode: 150-逆波兰表达式求值
		根据逆波兰表示法，求表达式的值。
		有效的算符包括+, -, *, /。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
	说明：
		整数除法只保留整数部分。
		给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

	逆波兰表达式：是一种后缀表达式，所谓后缀就是指算符写在后面。
	平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
	该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。

	逆波兰表达式主要有以下两个优点：
		去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
		适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。

	demo1:
		输入：tokens = ["2","1","+","3","*"]
		输出：9
		解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

	demo2:
		输入：tokens = ["4","13","5","/","+"]
		输出：6
		解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6


	思路: 使用栈LIFO特性
		1. 定义一个栈，用于储存运算数
		2. 遇到数字则入栈，遇到算符则取出栈顶两个数字进行计算(次栈顶元素 operate 栈顶元素 )，并将结果压入栈内

*/
package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	numStack := []int{}
	// 遍历切片
	for _, token := range tokens {
		// 尝试转换为整型
		val, err := strconv.Atoi(token)
		if err == nil {
			// 整型转换成功，说明是数字，则入栈
			numStack = append(numStack, val)
		} else {
			// 转换失败，为操作符，需要进行运算
			// 弹出栈顶、次栈顶元素
			num1, num2 := numStack[len(numStack)-1], numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-2]
			// 次栈顶元素 operate 栈顶元素
			switch token {
			case "+":
				numStack = append(numStack, num2+num1)
			case "-":
				numStack = append(numStack, num2-num1)
			case "*":
				numStack = append(numStack, num2*num1)
			case "/":
				numStack = append(numStack, num2/num1)
			}
		}
	}

	return numStack[0]
}
