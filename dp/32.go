package main

import "fmt"

// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

// 示例 1:

// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"

//状态转移方程
// ()()对应公式 dp[i]=dp[i-2]+2

//如果不是成对出现就取值 上一次出现的值  判断极值下标必须大于0,而且下标为左括号
func longestValidParentheses(s string) int {
	maxArea := 0
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxArea = max(maxArea, dp[i])
	}
	fmt.Println(dp)
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	longestValidParentheses(")()())")
}
