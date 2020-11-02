package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func twoSum(nums []int, target int) []int {
	mc := make(map[int]int)
	for index, value := range nums {
		if p, ok := mc[target-value]; ok {
			return []int{p, index}
		} else {
			mc[value] = index
		}
	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 5))
}
