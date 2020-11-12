package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var visited []int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	visited = append(visited, root.Val)
	if root.Left != nil {
		dfs(root.Left)
	}

	if root.Right != nil {
		dfs(root.Right)
	}
}

var visitedBfs []int

func bfs(stack []*TreeNode) {
	if len(stack) == 0 {
		return
	}
	visitedBfs = append(visitedBfs, stack[0].Val)

	if stack[0].Left != nil {
		stack = append(stack, stack[0].Left)
	}

	if stack[0].Right != nil {
		stack = append(stack, stack[0].Right)
	}

	stack = stack[1:]
	bfs(stack)
}

func main() {
	root1 := &TreeNode{Val: 1}
	root2 := &TreeNode{Val: 2}
	root3 := &TreeNode{Val: 3}
	root4 := &TreeNode{Val: 4}
	root5 := &TreeNode{Val: 5}
	root6 := &TreeNode{Val: 6}
	root7 := &TreeNode{Val: 7}

	root1.Left = root2
	root1.Right = root3

	root2.Left = root4
	root2.Right = root5

	root3.Left = root6
	root3.Right = root7
	// dfs(root1)

	var stack []*TreeNode
	stack = append(stack, root1)
	bfs(stack)
	fmt.Println(visitedBfs)
}
