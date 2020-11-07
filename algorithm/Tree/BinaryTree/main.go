package main

import (
	"fmt"
	"math"
)

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

//迭代实现前序遍历
//如何思考用递归解决问题

func RecurtionPrevTreeNode(root *TreeNode) {
	stack := []*TreeNode{}
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

//中序遍历 左-根-右
func MiddleTreeNode(tree *TreeNode) {
	if tree != nil {
		MiddleTreeNode(tree.Left)
		fmt.Println(tree.Val)
		MiddleTreeNode(tree.Right)
	}

}
func RecurtionMiddleTreeNode(root *TreeNode) {
	stack := []*TreeNode{}
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

//后序遍历 左-右-根
func AfterTreeNode(tree *TreeNode) {
	if tree != nil {
		AfterTreeNode(tree.Left)
		AfterTreeNode(tree.Right)
		fmt.Println(tree.Val)
	}
}

//后序遍历 左-右-根
//利用前序遍历  root-right-left 然后反转 好牛逼的想法
func RecurtionAfterTreeNode(root *TreeNode) {
	stack := []*TreeNode{}
	var result []int

	for root != nil || len(stack) > 0 {
		if root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Right
		} else {
			node := stack[len(stack)-1]
			root = node.Left
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Print(result)
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

//递归前序遍历left-root-right
func RecurtionMiddleTreeNode2(root *TreeNode) {
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

//翻转二叉树
//左子树与右子数相互对调
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

//验证二叉搜索树
// 假设一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。

// 所有左子树和右子树自身必须也是二叉搜索树。

//最牛逼的思路是中序遍历
//递归实现
func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)
}

func validBST(root *TreeNode, min, max int) bool {
	//递归结束条件
	if root == nil {
		return true
	}
	// 判断节点的值是不是在区间呢，不是的话就false结束
	if root.Val <= min || root.Val >= max {
		return false
	}

	return validBST(root.Left, min, root.Val) && validBST(root.Right, root.Val, max)
}

// 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 二叉树的最小深度
// 给定一个二叉树，找出其最小深度。
// 叶子节点到根节点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	} else {
		return min(left, right) + 1
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 236. 二叉树的最近公共祖先

//示例 1:

// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3
// 解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//边界值
	if root == nil || root == p || root == q {
		return root
	}
	//处理逻辑

	//进入下一层
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

//从前序与中序遍历序列构造二叉树
//反推树的结构

func buildTree(preorder []int, inorder []int) *TreeNode {
	var preLength, inLength int
	var c func(preLength, inLength, inEnd int, preorder []int, inorder []int) *TreeNode

	c = func(preLength, inLength, inEnd int, preorder []int, inorder []int) *TreeNode {
		if preLength > len(preorder)-1 || inLength > inEnd {
			return nil
		}
		root := &TreeNode{Val: preorder[preLength]}
		inIndex := 0 // Index of current root in inorder
		for i := inLength; i <= inEnd; i++ {
			if inorder[i] == root.Val {
				inIndex = i
			}
		}
		root.Left = c(preLength+1, inLength, inIndex-1, preorder, inorder)
		root.Right = c(preLength+inIndex-inLength+1, inIndex+1, inEnd, preorder, inorder)

		return root
	}

	return c(preLength, inLength, len(inorder)-1, preorder, inorder)
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
	fmt.Println(maxDepth(root1))
}
