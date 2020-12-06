package main

// 输入: [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

// 应该用BFS,但是这个如果用贪心法的思想去解题，应该执行更快
func jump(nums []int) int {
	end := 0
	steps := 0
	maxPosition := 0

	for i := 0; i < len(nums)-1; i++ {
		//每次都跳最大的跳跃式数
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
