package main

//移动0问题

//注意  必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
func moveZeroes(nums []int) {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
