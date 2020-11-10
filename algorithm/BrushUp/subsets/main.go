package main

import "fmt"

//78. 子集
// 示例:
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

// 说明：解集不能包含重复的子集。
// 输入: nums = [1,2,3]
// 输出:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

// [[],[3],[1,3],[1,2,3],[1,3],[3],[2,3],[3]]
// [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

//一筹莫展
func subsets(nums []int) [][]int {
	result := new([][]int)
	issubsets([]int{}, nums, 0, result)
	return *result
}

func issubsets(buf []int, nums []int, start int, result *[][]int) {
	temp := make([]int, len(buf))
	copy(temp, buf)
	*result = append(*result, temp)
	for i := start; i < len(nums); i++ {
		buf = append(buf, nums[i])
		issubsets(buf, nums, i+1, result)
		buf = buf[:len(buf)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
