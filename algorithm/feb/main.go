package main

import (
	"fmt"
	"sort"
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

//滑动窗口
// 480. 滑动窗口中位数
// 中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

// 例如：

// [2,3,4]，中位数是 3
// [2,3]，中位数是 (2 + 3) / 2 = 2.5
// 给你一个数组 nums，有一个长度为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。

// 示例：

// 给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。

// 窗口位置                      中位数
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       1
//  1 [3  -1  -3] 5  3  6  7      -1
//  1  3 [-1  -3  5] 3  6  7      -1
//  1  3  -1 [-3  5  3] 6  7       3
//  1  3  -1  -3 [5  3  6] 7       5
//  1  3  -1  -3  5 [3  6  7]      6
//  因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。

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

func main() {
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
