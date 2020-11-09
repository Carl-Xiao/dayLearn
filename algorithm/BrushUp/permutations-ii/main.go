package main

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

//

// 示例 1：

// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]]

func permuteUnique(nums []int) [][]int {
	result := &[][]int{}
	_permuteUnique(0, nums, result)
	return *result

}

func _permuteUnique(start int, nums []int, results *[][]int) {
	if start == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*results = append(*results, tmp)
		return
	}
	used := make(map[int]bool)

	for i := start; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		used[nums[i]] = true
		nums[i], nums[start] = nums[start], nums[i]
		_permuteUnique(start+1, nums, results)
		nums[i], nums[start] = nums[start], nums[i]
		used[nums[i]] = false
	}
}

func main() {

}
