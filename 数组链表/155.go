package main

import "fmt"

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。

//常数时间最小那就是第一个元素或者最后一个元素才对
type MinStack struct {
	stack    []int
	minstack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minstack: []int{},
	}

}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minstack) == 0 {
		this.minstack = append(this.minstack, x)
	} else {
		temp := this.minstack[len(this.minstack)-1]
		if temp > x {
			this.minstack = append(this.minstack, x)
		} else {
			this.minstack = append(this.minstack, temp)
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minstack = this.minstack[:len(this.minstack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	fmt.Println(len(this.minstack))
	return this.minstack[len(this.minstack)-1]
}
