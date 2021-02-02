package main

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
