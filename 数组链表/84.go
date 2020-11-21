package main

import "fmt"

//单调栈
//这个思想怎么说呢，也算比较好理解，但是不容易想到会这么处理逻辑
//栈里面存下标，才方便取数据算长度

// 示例:

// 输入: [2,1,5,6,2,3]
// 输出: 10
func largestRectangleArea(heights []int) int {
	var stack []int
	result := 0
	var width int
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			result = max(result, height*width)
		}
		stack = append(stack, i)
	}
	n := len(heights)
	fmt.Println(stack)
	for len(stack) > 0 {
		height := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			width = n
		} else {
			width = n - stack[len(stack)-1] - 1
		}

		result = max(result, heights[height]*width)
	}
	fmt.Println("=========")
	fmt.Println(result)
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
