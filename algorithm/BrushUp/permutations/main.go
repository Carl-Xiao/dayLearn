package main

import "fmt"

// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

// 示例:

// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

func permute(nums []int) [][]int {
	result := &[][]int{}
	_permute(0, nums, result)
	return *result
}
func _permute(k int, nums []int, result *[][]int) {
	if len(nums) == k {
		temp := make([]int, k)
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	fmt.Println(nums)
	for i := k; i < len(nums); i++ {
		//先改
		nums[i], nums[k] = nums[k], nums[i]
		_permute(k+1, nums, result)
		//还原
		nums[i], nums[k] = nums[k], nums[i]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
