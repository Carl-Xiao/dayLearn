package main

import "math"

// 515. 在每个树行中找最大值

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//这就是bfs遍历吗
func largestValues(root *TreeNode) []int {
	var stack []*TreeNode
	result := new([]int)

	if root == nil {
		return *result
	}
	stack = append(stack, root)

	__largestValues(stack, result)
	return *result
}

func __largestValues(stack []*TreeNode, result *[]int) {
	for len(stack) > 0 {
		size := len(stack)
		temp := math.MinInt64
		for size > 0 {
			root := stack[0]
			stack = stack[1:]
			if root == nil {
				continue
			}
			temp = max(root.Val, temp)

			if root.Left != nil {
				stack = append(stack, root.Left)
			}
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			size--
		}
		*result = append(*result, temp)
	}

}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {

}
