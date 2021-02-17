package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

//ListNode 节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//排序链表
//要求 O(n log n) 时间复杂度和常数级空间复杂度下，

//归并排序
func sortList(head *ListNode) *ListNode {

	return nil
}

//倒数第几个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	node := head
	headL := head
	//当前节点的个数,如果知道这一步也可以省略
	i := 0
	for node != nil {
		node = node.Next
		i++
	}
	for j := 0; j < i-k; j++ {
		headL = headL.Next
	}
	return headL
}

//两两交换链表中的节点
//用递归处理节点
func swapPairs(head *ListNode) *ListNode {
	//终止条件
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = swapPairs(tmp.Next)
	tmp.Next = head
	return tmp
}

func medianSlidingWindow(nums []int, k int) (ans []float64) {
	//如果K为偶数,则不存在中位数
	knums := make([]int, k)
	copy(knums, nums[:k])
	sort.Ints(knums)
	ans = append(ans, midle(knums, k))

	for i := k; i < len(nums); i++ {
		id := sort.SearchInts(knums, nums[i-k]) //窗口左移动 左边的元素移除
		knums[id] = nums[i]                     //窗口右边元素加入列表
		//特殊情况如果替换的第一个元素
		for id < k-1 && knums[id] > knums[id+1] {
			knums[id], knums[id+1] = knums[id+1], knums[id]
			id++
		}
		//替换的其他元素
		for id > 0 && knums[id] < knums[id-1] {
			knums[id], knums[id-1] = knums[id-1], knums[id]
			id--
		}
		ans = append(ans, midle(knums, k))
	}
	return ans
}

//计算中位数
func midle(nums []int, k int) float64 {
	n1, n2 := 0, 0
	if k%2 == 0 {
		n1, n2 = nums[k/2-1], nums[k/2]
	} else {
		n1, n2 = nums[k/2], nums[k/2]
	}
	return float64(n1+n2) / float64(2)
}

// 给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

//

// 示例：

// 输入：[1,12,-5,-6,50,3], k = 4
// 输出：12.75
// 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75

func findMaxAverage(nums []int, k int) (number float64) {
	kums := make([]int, k)
	copy(kums, nums[:k])
	var sum int
	for _, v := range kums {
		sum += v
	}
	number = calc(sum, k)
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - kums[0]
		kums = append(kums, nums[i])
		kums = kums[1:]
		number = max(number, calc(sum, k))
	}
	return
}
func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func calc(total int, k int) float64 {
	return float64(total) / float64(k)
}

// 给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

// 你可以返回满足此条件的任何数组作为答案。

//

// 示例：

// 输入：[3,1,2,4]
// 输出：[2,4,3,1]
// 输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。

func sortArrayByParity(A []int) (result []int) {
	var odd, even []int
	for _, v := range A {
		if v%2 == 0 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
		}
	}
	result = append(result, odd...)
	result = append(result, even...)
	return
}

// 给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

// 对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

// 你可以返回任何满足上述条件的数组作为答案。

// 示例：
// 输入：[4,2,5,7]
// 输出：[4,5,2,7]
// 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。

func sortArrayByParityII(A []int) (result []int) {
	result = make([]int, len(A))
	odd, even := 0, 1
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			result[odd] = A[i]
			odd += 2
		} else {
			result[even] = A[i]
			even += 2
		}
	}
	return
}

// 示例 1：
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]

//有序数的平方
func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	//排序没有利用升序的性能
	sort.Ints(nums)
	return nums
}

//利用双指针处理
// func sortedSquares(a []int) []int {
// 	n := len(a)
// 	ans := make([]int, n)
// 	i, j := 0, n-1
// 	for pos := n - 1; pos >= 0; pos-- {
// 		if v, w := a[i]*a[i], a[j]*a[j]; v > w {
// 			ans[pos] = v
// 			i++
// 		} else {
// 			ans[pos] = w
// 			j--
// 		}
// 	}
// 	return ans
// }

// 符合下列属性的数组 arr 称为 山脉数组 ：
// arr.length >= 3
// 存在 i（0 < i < arr.length - 1）使得：
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
// 给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。

//双指针
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l+1] > arr[l] {
			l++
		}
		if arr[r] < arr[r-1] {
			r--
		}
	}
	return l
}

// 给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。

// （例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）

// 返回 A 中好子数组的数目。

// 示例 1：

// 输入：A = [1,2,1,2,3], K = 2
// 输出：7
// 解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].

//滑动窗口解题
func subarraysWithKDistinct(A []int, K int) int {
	l := 0
	n := len(A)
	mac := make(map[int]int)
	distinct := 0
	res := 0
	result := 1
	for r := 0; r < n; r++ {
		if mac[A[r]] == 0 {
			distinct++
		}
		mac[A[r]]++

		for mac[A[l]] > 1 || distinct > K {
			if distinct > K {
				result = 1
				distinct--
			} else {
				result++
			}
			mac[A[l]]--
			l++
		}
		//判断map的size==K
		if distinct == K {
			res += result
		}
	}
	return res
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] {
			l++
		} else if nums[mid] == nums[r] {
			r--
		} else if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}

//
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			l = i + 1
			break
		}
	}
	return nums[l]
}

// 示例 1：
// 输入: [1,3,5]
// 输出: 1

// 示例 2：

// 输入: [2,2,2,0,1]
// 输出: 0
func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	minresult := math.MaxInt64
	for l <= r {
		if nums[l] < nums[r] {
			minresult = min(minresult, nums[l])
			l++
		} else {
			minresult = min(minresult, nums[r])
			r--
		}
	}
	return minresult
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func searchMatrix(matrix [][]int, target int) bool {
	var nums []int
	for i := 0; i < len(matrix); i++ {
		n := len(matrix[i])
		fmt.Println(matrix[i])
		if matrix[i][n-1] >= target && matrix[i][0] <= target {
			nums = matrix[i]
			break
		}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return false

}

//括号生成
func generateParenthesis(n int) []string {
	res := new([]string)
	_generateParent(n, n, "", res)
	return *res
}

func _generateParent(l, r int, s string, result *[]string) {
	if l == 0 || r == 0 {
		*result = append(*result, s)
		return
	}
	if l > 0 {
		_generateParent(l-1, r, s+"(", result)
	}
	if r > l {
		_generateParent(l, r-1, s+")", result)
	}
}

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	cols := map[int]bool{}
	pie := map[int]bool{}
	na := map[int]bool{}
	res := [][]string{}
	helper(0, bd, &res, n, &cols, &pie, &na)
	return nil
}
func helper(r int, bd [][]string, res *[][]string, n int, cols, diag1, diag2 *map[int]bool) {
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}
	for c := 0; c < n; c++ {
		if !(*cols)[c] && !(*diag1)[r+c] && !(*diag2)[r-c] {
			bd[r][c] = "Q"
			(*cols)[c] = true
			(*diag1)[r+c] = true
			(*diag2)[r-c] = true
			helper(r+1, bd, res, n, cols, diag1, diag2)
			bd[r][c] = "."
			(*cols)[c] = false
			(*diag1)[r+c] = false
			(*diag2)[r-c] = false
		}
	}
}

//求N皇后的解题数量
func totalNQueens(n int) (ans int) {
	columns := make([]bool, n)        // 列上是否有皇后
	diagonals1 := make([]bool, 2*n-1) // 左上到右下是否有皇后
	diagonals2 := make([]bool, 2*n-1) // 右上到左下是否有皇后
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}
		for col, hasQueen := range columns {
			d1, d2 := row+n-1-col, row+col
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			backtrack(row + 1)
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	backtrack(0)
	return
}

//搜索二维矩阵 II
func searchMatrixII(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
			continue
		}
		if matrix[row][col] < target {
			row++
			continue
		}
	}
	return false
}

//合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	//只保留一个节点
	for len(lists) > 1 {
		lists = append(lists, merge(lists[0], lists[1]))
		lists = lists[2:]
	}
	return lists[0]
}

func merge(one *ListNode, two *ListNode) *ListNode {
	root := new(ListNode)
	cur := root
	for one != nil && two != nil {
		if one.Val < two.Val {
			cur.Next = one
			one = one.Next
		} else {
			cur.Next = two
			two = two.Next
		}
		cur = cur.Next
	}
	if one != nil {
		cur.Next = one
	} else {
		cur.Next = two
	}
	return root.Next
}

func findKthLargest(nums []int, k int) int {
	return Partition(nums, k)
}

func Partition(nums []int, k int) int { //快排
	i, j := 0, len(nums)-1
	pivot := nums[0]
	for i < j {
		for i < j && nums[j] < pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		for i < j && nums[i] > pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	if i+1 == k {
		return pivot //   位子恰好就是K-1返回结果
	} else if i+1 < k {
		return Partition(nums[i+1:], k-i-1) //以pivot为断点，在后半部分数组中进行查找
	} else {
		return Partition(nums[:i], k) //在前半部分数组中进行查找
	}
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n, m := len(nums), len(nums[0])
	if n*m != r*c {
		return nums
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	//压缩成为一维数组
	for i := 0; i < n*m; i++ {
		ans[i/c][i%c] = nums[i/m][i%m]
	}
	return ans
}

func main() {
	fmt.Println(solveNQueens(4))
}
