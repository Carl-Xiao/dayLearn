package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历 根-左-右
func PrevTreeNode(tree *TreeNode) {
	if tree != nil {
		fmt.Println(tree.Val)
		PrevTreeNode(tree.Left)
		PrevTreeNode(tree.Right)
	}

}

//中序遍历 左-根-右
func MiddleTreeNode(tree *TreeNode) {
	if tree != nil {
		MiddleTreeNode(tree.Left)
		fmt.Println(tree.Val)
		MiddleTreeNode(tree.Right)
	}

}

//后序遍历 左-右-根
func AfterTreeNode(tree *TreeNode) {
	if tree != nil {
		AfterTreeNode(tree.Left)
		AfterTreeNode(tree.Right)
		fmt.Println(tree.Val)
	}
}

type Node struct {
	Val      int
	Children []*Node
}

//N叉树后序遍历
func postorder(root *Node) (result []int) {
	var nodeHandle func(root *Node)
	nodeHandle = func(root *Node) {
		if root != nil {
			for _, value := range root.Children {
				nodeHandle(value)
			}
			result = append(result, root.Val)
		}
	}
	nodeHandle(root)
	return
}

//N叉树前序遍历
func preorder(root *Node) (result []int) {
	var nodeHandle func(root *Node)
	nodeHandle = func(root *Node) {
		if root != nil {
			result = append(result, root.Val)
			for _, value := range root.Children {
				nodeHandle(value)
			}
		}
	}
	nodeHandle(root)
	return
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
