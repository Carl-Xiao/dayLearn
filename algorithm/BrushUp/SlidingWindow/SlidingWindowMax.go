// #### [239. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)

// 给定一个数组 *nums*，有一个大小为 *k* 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 *k* 个数字。滑动窗口每次只向右移动一位。

// 返回滑动窗口中的最大值。

//leetcode #239

// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
// 解释:

//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

package main

import "fmt"

func MaxSlidingWindow(nums []int, k int) []int {
	// 双端队列 队头出 队尾进
	var queue []int
	var result []int
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)

		if queue[0] == i-k {
			queue = queue[1:]
		}

		if i-k+1 >= 0 {
			// 队头出
			result = append(result, nums[queue[0]])
		}
	}
	return result
}

func main() {
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
