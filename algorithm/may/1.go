package main

import (
	"strconv"
	"strings"
)

//1 两数之和
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

// 你可以按任意顺序返回答案。
// 并返回它们的数组下标
func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for i, v := range nums {
		if j, ok := maps[target-v]; ok {
			return []int{i, j}
		} else {
			maps[v] = i
		}
	}
	return []int{}
}

//20
//有效的括号
//利用栈解决
func isValid(s string) bool {
	var stack []byte
	for _, v := range s {
		if v == ')' {
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if v == ']' {
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if v == '}' {
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, byte(v))
	}
	return true
}

//394 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
//
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

//146 最少使用缓存
//map加双向链表

//找不同
func findTheDifference(s string, t string) byte {
	tMap := make(map[byte]int)
	for _, v := range s {
		tMap[byte(v)]++
	}
	for _, v := range t {
		tMap[byte(v)]--
	}
	for i, v := range tMap {
		if v > 0 {
			return i
		}
	}
	var b byte
	return b
}

//找单词规律
func wordPattern(pattern string, s string) bool {
	return true
}

//387 字符串第一个唯一的字符
func firstUniqChar(s string) int {

	return 0
}
