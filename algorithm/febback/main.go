package main

import "fmt"

//TreeNode 树节点
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

// [1,1,0,1,1,1]
// 输出: 3
// 解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
func findMaxConsecutiveOnes(nums []int) int {
	cnt := 0
	maxCnt := 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if maxCnt < cnt {
			maxCnt = cnt
		}
	}
	return maxCnt
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return cost[0]
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	nums := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		nums[i] = min(nums[i-1]+cost[i-1], nums[i-2]+cost[i-2])
	}
	return nums[len(cost)]
}

//爬楼梯问题
func climbStairs(n int) int {
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		b = a + b
		a, b = b, a
	}
	return a
}

//dp方程
//dp[i][j]=dp[i-1][j]+dp[i][j-1]
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//路径和最小,每一次的取值最小
//dp[i][j]=min(dp[i-1][j]+grid[i-1][j],dp[i][j-1]+grid[i][j-1])
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}

//最大子序和
//dp[i]=max(dp[i-1]+nums[i],nums[i])
// -2,1,-3,4,-1,2,1,-5,4
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//乘积最大子数组
//正负数的不确定,需要保存一个最小的数
func maxProduct(nums []int) (ans int) {
	maxF, minF := 1, 1
	ans = nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			minF, maxF = maxF, minF
		}
		maxF = max(maxF*nums[i], nums[i])
		minF = min(minF*nums[i], nums[i])
		ans = max(maxF, ans)
	}
	return
}

//买卖股票问题
//0 是卖出
//1 是买入

//dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
//昨天卖出/昨天买进消耗的钱与今天卖出转的利润相加
//dp[i][1]=max(dp[i-1][1],-prices[i])
//昨天买进/今天买进
//一次买进/卖出
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

//多次买进/卖出
//0 是卖出
//1 是买入
//dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
//dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
func maxProfitII(prices []int) int {
	n := len(prices)
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

//给定次数的买进/卖出
//0 是卖出
//1 是买入
//dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
//dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
func maxProfitIII(prices []int) int {
	return 0
}

//零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
// 你可以认为每种硬币的数量是无限的。
//使用动态规划
//j代表不同种类的硬币

//dp[i]=min(dp[i-conins[j]]+1，dp[i])
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 518. 零钱兑换 II
//计算组合数的个数
//f(n) = f(n-coins[0]) +  f(n-coins[1]) + ...
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i < amount+1; i++ {
			if i < coin {
				continue
			}
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

//
// 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

// 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

func maxEnvelopes(envelopes [][]int) int {

	return 0
}

// 300. 最长递增子序列
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}


func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
