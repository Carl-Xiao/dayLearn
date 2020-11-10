package main

import "fmt"

// 22. 括号生成
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
	_generateParenthesis("", n, 0, 0, result)
	return *result
}

func _generateParenthesis(buf string, n, left, right int, result *[]string) {
	if left == n && right == n {
		*result = append(*result, buf)
		return
	}
	//处理逻辑
	if left < n {
		_generateParenthesis(buf+"(", n, left+1, right, result)
	}
	if left > right {
		_generateParenthesis(buf+")", n, left, right+1, result)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
