package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	//设置哨兵节点
	result := root
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			root.Next = l1
			l1 = l1.Next
		} else {
			root.Next = l2
			l2 = l2.Next
		}
		root = root.Next
	}
	//处理残余情况
	if l1 != nil {
		root.Next = l1
	}
	if l2 != nil {
		root.Next = l2
	}

	return result.Next
}

//环形链表
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

//206反转链表 不可取
func reverseList(head *ListNode) *ListNode {
	root := &ListNode{}
	reuslt := root
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	for i := 0; i < len(nums); i++ {
		root.Next = &ListNode{Val: nums[i]}
		root = root.Next
	}
	return reuslt.Next
}

//206反转链表 参考官方
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode

	return prev
}

//反转从位置m到n的链表
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	prev := &ListNode{Next: head}
	result := prev

	for i := 1; i <= m-1; i++ {
		result = result.Next
	}
	current := result.Next
	//反转
	var pre *ListNode
	for i := m; i <= n; i++ {
		tempNext := current.Next
		current.Next = pre
		pre = current
		current = tempNext
	}
	//拼接n的部分
	result.Next.Next = current
	//拼接m-n的部分
	result.Next = pre

	return prev.Next
}

func max(x, y int) {

}
