package main

import "fmt"

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 输入：n = 3
// 输出：[
//        "((()))",
//        "(()())",
//        "(())()",
//        "()(())",
//        "()()()"
//      ]
//难点 这玩意要验证括号成对的正确性
//1 找重复性代码
//2 验证代码的合理性

//左括号阔以随意出现,右括号必须跟在左括号后边
//右括号的个数必须小于左括号

func generateParenthesis(n int) []string {
	result := new([]string)
	parnethesis(0, 0, n, "", result)
	return *result
}

func parnethesis(left, right, max int, s string, result *[]string) {
	if left == max && right == max {
		*result = append(*result, s)
		return
	}
	if left < max {
		parnethesis(left+1, right, max, s+"(", result)
	}
	if left > right {
		parnethesis(left, right+1, max, s+")", result)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
