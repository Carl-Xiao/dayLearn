package main

import "fmt"

//#941 有效的山脉数组

// 示例 1：

// 输入：[2,1]
// 输出：false
// 示例 2：

// 输入：[3,5,5]
// 输出：false
// 示例 3：

// 输入：[0,3,2,1]
// 输出：true

//用折半的方法，双指针找如果有一个不满足就返回

//国际版好思路 https://leetcode.com/problems/valid-mountain-array/discuss/194900/C%2B%2BJavaPython-Climb-Mountain

//一个人从左边爬山
//一个人从右边爬山
//最高点是同一个位置则代表是有效的山脉数组
func validMountainArray(A []int) bool {
	var n, i, j = len(A), 0, len(A) - 1
	for i+1 < n && A[i] < A[i+1] {
		i++
	}
	for j > 0 && A[j] < A[j-1] {
		j--
	}
	//容易忽略的点  不应该包含起点在内
	return i > 0 && j < n-1 && i == j
}

func main() {
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}
