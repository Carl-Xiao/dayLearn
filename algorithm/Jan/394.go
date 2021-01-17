package main

import (
	"strconv"
	"strings"
)

// 394. 字符串解码
func decodeString(s string) string {
	//当前使用两个栈维护
	numStack := []int{}
	strStack := []string{}
	num := 0
	result := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			strStack = append(strStack, result)
			result = ""
			numStack = append(numStack, num)
			num = 0
		} else if char == ']' {
			//弹出栈
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			//弹出栈
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			result = string(str) + strings.Repeat(result, count)
		} else {
			result += string(char)
		}
	}
	return result
}
