package main

import (
	"sort"
)

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	return nums[len(nums)/2]
}
