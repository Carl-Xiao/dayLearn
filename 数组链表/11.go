package main

//使用双指针处理数据
func maxArea(height []int) int {
	var maxArea int
	left := 0
	right := len(height) - 1
	for left <= right {
		var temp int
		if height[left] >= height[right] {
			temp = height[right]
			right--
		} else {
			temp = height[left]
			left++
		}
		maxArea = max(maxArea, temp*(right-left+1))
		// fmt.Println(maxArea, right, left)
	}

	return maxArea
}
