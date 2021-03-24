package main

import (
	"fmt"
	"math"
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

func main() {
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
}
