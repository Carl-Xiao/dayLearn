package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//104. 二叉树的最大深度
//使用dfs
func maxDepth(root *TreeNode) int {
	return _maxDepth(1, root)
}

func _maxDepth(depth int, root *TreeNode) int {
	if root == nil {
		return depth
	}
	return max(_maxDepth(depth+1, root.Left), _maxDepth(depth+1, root.Right))
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//112. 路径总和
//根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum
//终止条件是 到叶子节点未空或者找到节点数据
//dfs遍历
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := new([][]int)
	dfsSum(root, targetSum, []int{}, result)
	return *result
}

func dfsSum(root *TreeNode, targetSum int, t []int, result *[][]int) {
	if root == nil {
		return
	}
	targetSum -= root.Val
	t = append(t, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		*result = append(*result, append([]int(nil), t...))
		return
	}
	dfsSum(root.Left, targetSum, t, result)
	dfsSum(root.Right, targetSum, t, result)
	t = t[:len(t)-1]
}

// 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

// 示例:

// X X X X
// X O O X
// X X O X
// X O X X

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	if len(board[0]) == 0 {
		return
	}
	r, c := len(board), len(board[0])
	for i := 0; i < r; i++ {
		dfsSolve(board, i, 0)
		dfsSolve(board, i, c-1)
	}
	for i := 1; i < c-1; i++ {
		dfsSolve(board, 0, i)
		dfsSolve(board, r-1, i)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
func dfsSolve(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	dfsSolve(board, i-1, j)
	dfsSolve(board, i+1, j)
	dfsSolve(board, i, j-1)
	dfsSolve(board, i, j+1)
}

//计算岛屿的数量
//使用并查集
//或者使用dfs
func numIslands(grid [][]byte) int {
	var nums int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				nums++
				dfsLand(grid, len(grid), len(grid[i]), i, j)
			}
		}
	}
	return nums
}

//
func dfsLand(grid [][]byte, r, l, i, j int) {
	if i < 0 || j < 0 || i >= r || j >= l || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfsLand(grid, r, l, i+1, j)
	dfsLand(grid, r, l, i-1, j)
	dfsLand(grid, r, l, i, j-1)
	dfsLand(grid, r, l, i, j+1)
}

func maxAreaOfIsland(grid [][]int) int {
	var area int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				area = max(area, dfsAreaLand(grid, len(grid), len(grid[i]), i, j))
			}
		}
	}
	return area
}

func dfsAreaLand(grid [][]int, r, l, i, j int) int {
	if i < 0 || j < 0 || i >= r || j >= l || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfsAreaLand(grid, r, l, i+1, j) + dfsAreaLand(grid, r, l, i-1, j) + dfsAreaLand(grid, r, l, i, j-1) + dfsAreaLand(grid, r, l, i, j+1)
}

//分配硬币
func distributeCoins(root *TreeNode) int {
	res := 0
	dfsDistribute(root, &res)
	return res
}

func dfsDistribute(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	//后序遍历
	l, r := dfsDistribute(root.Left, res), dfsDistribute(root.Right, res)
	*res += abs(l) + abs(r)
	return l + r + root.Val - 1
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func main() {

}
