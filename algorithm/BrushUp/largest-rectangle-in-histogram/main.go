package main

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x

}

func largestRectangleArea(heights []int) int {
	var width, height int
	var result int
	var stack []int
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			height = heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			fmt.Println(stack)
			fmt.Println(width, height)
			result = max(result, width*height)
		}
		stack = append(stack, i)
	}
	fmt.Println("222")
	//stack还有值,则还需要额外处理
	n := len(heights)
	for len(stack) > 0 {
		height = heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			width = n
		} else {
			width = n - stack[len(stack)-1] - 1
		}
		fmt.Println(width, height)
		result = max(result, width*height)
	}
	return result
}

func main() {
	fmt.Println(largestRectangleArea([]int{3, 2, 1}))
}
