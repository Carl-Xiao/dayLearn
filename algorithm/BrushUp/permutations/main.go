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

func _permute(start int, nums []int, results *[][]int) {
	if start == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*results = append(*results, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		_permute(start+1, nums, results)
		nums[i], nums[start] = nums[start], nums[i]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
