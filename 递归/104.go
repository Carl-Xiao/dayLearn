package main

// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

func maxDepth(root *TreeNode) int {
	//边界值
	if root == nil {
		return 0
	}
	//进入多少层这个是明确知道的
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
