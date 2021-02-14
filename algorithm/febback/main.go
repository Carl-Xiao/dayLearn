package main

import (
	"fmt"
)

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

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfsWord(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfsWord(board [][]byte, word string, i, j, k int) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}

	c := board[i][j]
	board[i][j] = ' '

	if dfsWord(board, word, i+1, j, k+1) {
		return true
	}

	if dfsWord(board, word, i-1, j, k+1) {
		return true
	}

	if dfsWord(board, word, i, j+1, k+1) {
		return true
	}

	if dfsWord(board, word, i, j-1, k+1) {
		return true
	}
	board[i][j] = c
	return false
}

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] == i+1 { // 当前元素出现在它该出现的位置，无需交换
			i++
			continue
		}
		idealIdx := nums[i] - 1        // idealIdx：当前元素应该出现的位置
		if nums[i] == nums[idealIdx] { // 当前元素=它理应出现的位置上的现有元素，说明重复了
			i++
			continue
		}
		nums[idealIdx], nums[i] = nums[i], nums[idealIdx] // 不重复，进行交换
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 { // 值与索引 不对应
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func minOperations(s string) int {
	n := len(s)
	bs := []byte(s)
	a, b := make([]byte, len(bs)), make([]byte, len(bs))
	copy(a, bs)
	copy(b, bs)
	var x, y int
	//01
	for i := 1; i < len(a); {
		if a[i] == '1' {
			if a[i-1] == '0' {

			} else {
				a[i-1] = '0'
				x++
			}
		} else {
			a[i] = '1'
			x++
			if a[i-1] != '0' {
				a[i-1] = '0'
				x = x + 1
			}
		}
		i = i + 2
	}
	if n > 2 && a[n-1] == a[n-2] {
		x++
	}
	//10
	for i := 1; i < len(b); {
		if b[i] == '0' {
			if b[i-1] == '1' {

			} else {
				b[i-1] = '1'
				y++
			}
		} else {
			b[i] = '0'
			y++
			if b[i-1] != '1' {
				b[i-1] = '1'
				y = y + 1
			}
		}
		i = i + 2
	}
	if n > 2 && b[n-1] == b[n-2] {
		y++
	}
	return min(x, y)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countHomogenous(s string) int {
	r := 0
	mac := make(map[byte]int)
	var stack []byte
	var total int
	for r < len(s) {
		if _, ok := mac[s[r]]; !ok {
			stack = append(stack, s[r])
		}
		mac[s[r]]++
		for len(mac) == 2 {
			v := mac[stack[0]]
			total += calc(v)
			delete(mac, stack[0])
			stack = stack[1:]
		}
		r++
	}
	for _, v := range mac {
		total += calc(v)
	}
	return total % 1000000007
}

func calc(v int) int {
	var total int
	for i := 1; i <= v; i++ {
		total += i
	}
	return total
}

func check(nums []int, cost, maxOperations int) bool {
	var ans int
	for _, cur := range nums {
		if cur%cost == 0 {
			ans += cur/cost - 1
		} else {
			ans += cur / cost
		}
	}
	return ans <= maxOperations
}

func minimumSize(nums []int, maxOperations int) int {
	l := 1
	r := 1000000000
	ret := 0
	for l <= r {
		mid := (l + r) / 2
		if check(nums, mid, maxOperations) {
			r = mid - 1
			ret = mid
		} else {
			l = mid + 1
		}
	}
	return ret
}

func main() {
	// 10010100
	fmt.Println(countHomogenous("zzzzz"))
}
