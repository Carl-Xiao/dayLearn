package main

import "fmt"

// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

// 返回滑动窗口中的最大值。

//利用队列
func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	var result []int
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		//移除K个元素
		//例如 7 2 4 当4加进来 变成 7 4 但是因为需要队头出队列把7移除
		if queue[0] == i-k {
			queue = queue[1:]
		}
		//取值
		if i-k+1 >= 0 {
			result = append(result, nums[queue[0]])
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	maxSlidingWindow([]int{7, 2, 4}, 2)
}
