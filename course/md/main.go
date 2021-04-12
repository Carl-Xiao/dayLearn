package main

import (
	"fmt"
	"math"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Left, root.Val, max)
}

type Node struct {
	Value    int
	NextNode *Node
}

func reverse(headNode *Node) *Node {
	var result *Node
	for headNode != nil {
		next := headNode.NextNode
		headNode.NextNode = result
		result = headNode
		headNode = next
	}
	return result
}

type user struct {
}

type send interface {
	send()
}

// func (u user) send() {
// 	fmt.Println("this a send")
// }

func (u *user) send() {
	fmt.Println("this a send")
}

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
