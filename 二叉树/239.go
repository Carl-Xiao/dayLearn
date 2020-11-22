package main

//滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	var result []int
	for i := 0; i < len(nums); i++ {
		//1 维护一个队列
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		//2 筛选合适长度
		if queue[0] == i-k {
			queue = queue[1:]
		}
		//3 加入result
		if i-k+1 >= 0 {
			result = append(result, nums[queue[0]])

		}
	}
	return result
}
