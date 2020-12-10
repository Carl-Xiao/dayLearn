package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
//前序
func preSort(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preSort(root.Left)
	preSort(root.Right)
}

//中序

func inorderTraversal(root *TreeNode) []int {
	result := new([]int)
	midSort(root, result)
	return *result
}

func midSort(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	midSort(root.Left, result)
	*result = append(*result, root.Val)
	midSort(root.Right, result)
}

//后序
func afterSort(root *TreeNode) {
	if root == nil {
		return
	}
	preSort(root.Left)
	preSort(root.Right)
	fmt.Println(root.Val)
}

//非递归

//手动创建递归栈
func RecursionPreSort(root *TreeNode) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			fmt.Println(root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
}

func RecursionMidSort(root *TreeNode) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			fmt.Println(node.Val)
			root = node.Right
			stack = stack[:len(stack)-1]
		}
	}
}

// func main() {
// 	root1 := &TreeNode{Val: 1}
// 	root2 := &TreeNode{Val: 2}
// 	root3 := &TreeNode{Val: 3}

// 	// root4 := &TreeNode{Val: 4}
// 	// root5 := &TreeNode{Val: 5}
// 	// root6 := &TreeNode{Val: 6}
// 	// root7 := &TreeNode{Val: 7}

// 	root1.Left = root2
// 	root1.Right = root3

// 	// root2.Left = root4
// 	// root2.Right = root5

// 	// root3.Left = root6
// 	// root3.Right = root7
// 	RecursionPreSort(root1)
// }
