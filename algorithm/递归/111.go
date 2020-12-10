package main

// 给定一个二叉树，找出其最小深度。

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	//如果的极端情况,单链表情况
	if left == 0 || right == 0 {
		return left + right + 1
	}

	return min(left, right) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
