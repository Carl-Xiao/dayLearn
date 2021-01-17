package main

//普通解法
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

//进阶解法
func twoSumSecond(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for index, value := range nums {
		if i, ok := numsMap[target-value]; ok {
			return []int{index, i}
		}
		numsMap[value] = index

	}
	return []int{}
}
