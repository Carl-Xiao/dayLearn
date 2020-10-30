package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrevTreeNode(tree *TreeNode) {
	if tree != nil {
		fmt.Println(tree.Val)
		PrevTreeNode(tree.Left)
		PrevTreeNode(tree.Right)
	}

}

func MiddleTreeNode(tree *TreeNode) {
	if tree != nil {
		MiddleTreeNode(tree.Left)
		fmt.Println(tree.Val)
		MiddleTreeNode(tree.Right)
	}

}

func AfterTreeNode(tree *TreeNode) {
	if tree != nil {
		AfterTreeNode(tree.Left)
		AfterTreeNode(tree.Right)
		fmt.Println(tree.Val)
	}
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
	MiddleTreeNode(root1)
}
