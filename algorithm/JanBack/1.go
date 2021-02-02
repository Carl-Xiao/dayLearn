package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// 290. 单词规律

//输入: pattern = "abba", str = "dog cat cat dog"
// 输出: true
//
func wordPattern(pattern string, s string) bool {
	str2byte := map[string]byte{}
	byte2str := map[byte]string{}
	str := strings.Split(s, " ")
	if len(str) != len(pattern) {
		return false
	}
	for i := range pattern {
		v := byte(pattern[i])
		if str2byte[str[i]] > 0 && str2byte[str[i]] != v || byte2str[v] != "" && byte2str[v] != str[i] {
			return false
		}
		str2byte[str[i]] = v
		byte2str[v] = str[i]
	}
	return true
}

// 387. 字符串中的第一个唯一字符

// s = "leetcode"
// 返回 0

// s = "loveleetcode"
// 返回 2
//mac 处理

func firstUniqChar(s string) int {
	mac := make(map[byte]int)
	for i := range s {
		v := s[i] - 'a'
		mac[v]++
	}

	for i := range s {
		v := s[i] - 'a'
		if mac[v] == 1 {
			return i
		}
	}
	return -1
}

// 示例 1:

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	mac := make(map[byte]int)
	n := len(s)
	l, r, ans := 0, 0, 0
	for l < n {
		if r < n && mac[s[r]] == 0 {
			mac[s[r]] = 1
			r++
		} else {
			mac[s[l]] = 0
			l++
		}
		ans = max(ans, r-l)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//最小的覆盖子串

// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"

//滑动窗口，并且要校验整个子串出现的次数
func minWindow(s string, t string) string {
	mac := map[byte]int{}
	//值
	for i := range t {
		mac[t[i]]++
	}
	//类型
	windowType := len(mac)
	//最小值
	minlen, start := len(s), len(s)
	n := len(s)
	l, r := 0, 0
	for ; r < n; r++ {
		v := s[r]
		if result, ok := mac[v]; ok {
			mac[v] = result - 1
			if result-1 == 0 {
				windowType--
			}
		}
		for windowType == 0 {
			if r-l+1 < minlen {
				minlen = r - l + 1
				start = l
			}
			t := s[l]
			if result, ok := mac[t]; ok {
				mac[t] = result + 1
				if result+1 > 0 {
					windowType++
				}
			}
			l++
		}
	}
	if start == len(s) {
		return ""
	}
	return s[start : start+minlen]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//合并升序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	for l1 != nil && l2 != nil {
		v1 := l1.Val
		v2 := l2.Val
		if v1 < v2 {
			root.Next = l1
			l1 = l1.Next
		} else {
			root.Next = l2
			l2 = l2.Next
		}
		root = root.Next
	}
	if l1 != nil {
		root.Next = l1
	}

	if l2 != nil {
		root.Next = l2
	}
	return node.Next
}

//判断链表是否有环
func hasCycle(head *ListNode) bool {
	mac := map[int]int{}
	node := head
	for node != nil {
		if _, ok := mac[node.Val]; ok {
			return true
		}
		mac[node.Val] = 1
		node = node.Next
	}
	return false
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	var root *ListNode
	for head != nil {
		next := head.Next
		head.Next = root
		root = head
		head = next
	}
	return root
}

//反转链表m-n
//1 2 3 4 5
//1
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	prev := &ListNode{Next: head}
	reulst := prev
	for i := 1; i <= m-1; i++ {
		reulst = reulst.Next
	}
	current := reulst.Next
	ppNode := reulst.Next
	var reverseNode *ListNode
	for i := m; i <= n; i++ {
		next := current.Next
		current.Next = reverseNode
		reverseNode = current
		current = next
	}
	//  最后就变成了1 -> 2 <- 3 <- 4 -> 5 -> NULL
	ppNode.Next = current
	reulst.Next = reverseNode
	return prev.Next
}

type entry struct{ i, j, val int }
type hp []entry

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(entry)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type pair struct{ x, y int }

//上下左右值
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func swimInWater(grid [][]int) (ans int) {
	n := len(grid)
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[0][0] = true
	h := &hp{{0, 0, grid[0][0]}}
	for {
		e := heap.Pop(h).(entry)
		ans = max(ans, e.val)
		if e.i == n-1 && e.j == n-1 {
			return
		}
		for _, d := range dirs {
			if x, y := e.i+d.x, e.j+d.y; 0 <= x && x < n && 0 <= y && y < n && !vis[x][y] {
				vis[x][y] = true
				heap.Push(h, entry{x, y, grid[x][y]})
			}
		}
	}
}

//704. 二分查找
//有序
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		temp := nums[mid]
		if temp == target {
			return mid
		} else if temp > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

// 示例 1:

// 输入: 1->2->3->4->5->NULL, k = 2
// 输出: 4->5->1->2->3->NULL
// 解释:
// 向右旋转 1 步: 5->1->2->3->4->NULL
// 向右旋转 2 步: 4->5->1->2->3->NULL

// 61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	tmp := head
	//计算数量,这一步必不可少
	count := 1
	for tmp.Next != nil {
		tmp = tmp.Next
		count++
	}
	//如果刚好绕整圈,则代表位置不变
	k = k % count
	if k == 0 {
		return head
	}

	newHead, newTail := head, head
	for i := 0; i < count-k; i++ {
		newHead = newHead.Next
	}

	for i := 0; i < count-k-1; i++ {
		newTail = newTail.Next
	}

	tmp.Next = head
	newTail.Next = nil

	return newHead
}

// 424. 替换后的最长重复字符
// 给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

// 注意：字符串长度 和 k 不会超过 104。

// 示例 1：

// 输入：s = "ABAB", k = 2
// 输出：4
// 解释：用两个'A'替换为两个'B',反之亦然。

func characterReplacement(s string, k int) int {
	l, r := 0, 0
	if s == "" {
		return 0
	}
	historyCharMax := 0
	var mac [26]int
	for ; r < len(s); r++ {
		v := s[r] - 'A'
		mac[v]++
		historyCharMax = max(historyCharMax, mac[v])
		if r-l+1 > historyCharMax+k {
			mac[s[l]-'A']--
			l++
		}
	}
	return len(s) - l
}

// BANC
func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}
