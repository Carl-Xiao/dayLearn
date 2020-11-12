package main

// 102. 二叉树的层序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := new([]*TreeNode)
	result := new([][]int)
	*queue = append(*queue, root)
	_levelOrder(*queue, result)
	return *result
}

func _levelOrder(queue []*TreeNode, result *[][]int) {
	for len(queue) != 0 {
		var temp []int
		nums := len(queue)
		for i := 0; i < nums; i++ {
			if queue[0] == nil {
				queue = queue[1:]
				continue
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			temp = append(temp, queue[0].Val)
			queue = queue[1:]
		}
		if len(temp) > 0 {
			*result = append(*result, temp)
		}
	}
}

func main() {

}
