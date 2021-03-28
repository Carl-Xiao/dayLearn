package main

import (
	"fmt"
	"math"
	"strings"
)

//132 题目
// 示例 1：

// 输入：nums = [1,2,3,4]
// 输出：false
// 解释：序列中不存在 132 模式的子序列。

// 示例 2：

// 输入：nums = [3,1,4,2]
// 输出：true
// 解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。

// i j k
//使用单调递减栈处理
func find132pattern(nums []int) bool {
	n := len(nums)
	stack := []int{nums[n-1]}
	//维护最大值
	maxFn := math.MinInt64
	for i := n - 2; i >= 0; i-- {
		//再找到一个I就能满足条件
		if nums[i] < maxFn {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			maxFn = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if nums[i] > maxFn {
			stack = append(stack, nums[i])
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur: root,
	}
}

func (this *BSTIterator) Next() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	v, ok := <-c
	fmt.Println(v, ok)
}

//不同的整数
func numDifferentIntegers(word string) int {
	mac := make(map[string]int)
	words := []byte(word)
	for i, v := range words {
		if v >= 'a' && v <= 'z' {
			words[i] = ','
		}
	}
	for _, v := range strings.Split(string(words), ",") {
		if v == "" {
			continue
		}
		if len(v) == 1 {
			mac[v]++
			continue
		}
		index := strings.Index(v, "0")
		if index > 0 {
			mac[v]++
			continue
		}
		lastIndex := strings.LastIndex(v, "0")
		if index == lastIndex {
			mac[v[index+1:]]++
		} else {
			mac[v[lastIndex+1:]]++
		}
	}
	return len(mac)
}

// 5714. 替换字符串中的括号内容
func evaluate(s string, knowledge [][]string) string {
	mac := make(map[string]string)
	for _, v := range knowledge {
		mac[v[0]] = v[1]
	}
	result := ""
	tmp := ""
	flag := false
	words := []byte(s)
	for i := 0; i < len(words); i++ {
		if words[i] == '(' {
			flag = true
			tmp = ""
		} else if words[i] == ')' {
			flag = false
			if v, ok := mac[tmp]; ok {
				result += v
			} else {
				result += "?"
			}
		} else {
			if flag {
				tmp += string(words[i])
			} else {
				result += string(words[i])
			}
		}
	}
	return result
}

func main() {
	fmt.Println(evaluate("hi(name)", [][]string{
		{"a", "b"},
	}))
}
