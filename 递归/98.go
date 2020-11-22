package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 假设一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

//验证二叉搜索树 说白了就是比较左右子节点的大小问题

func isValidBST(root *TreeNode) bool {
	return _isValidBST(root, math.MinInt64, math.MaxInt64)
}

func _isValidBST(root *TreeNode, min, max int) bool {
	//临界条件
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	//左子树或者右子树
	return _isValidBST(root.Left, min, root.Val) && _isValidBST(root.Right, root.Val, max)
}
