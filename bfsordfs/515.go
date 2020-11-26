package main

import (
	"fmt"
	"math"
)

//TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//bfs 用queue不能用stack
func largestValues(root *TreeNode) []int {
	var queue []*TreeNode
	var result []int
	if root == nil {
		return []int{}
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		res := math.MinInt64

		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				res = max(res, node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			size--
		}
		fmt.Println(queue)
		result = append(result, res)

	}
	return result
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {

}
