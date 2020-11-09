package main

import "fmt"

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

// 示例:

// 输入: n = 4, k = 2
// 输出:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

func combine(n int, k int) [][]int {
	result := &[][]int{}
	_combine([]int{}, 1, n, k, result)
	return *result
}

func _combine(buf []int, begin, n, k int, result *[][]int) {
	if k == len(buf) {
		temp := make([]int, k)
		copy(temp, buf)
		*result = append(*result, temp)
		return
	}
	for i := begin; i <= n; i++ {
		buf = append(buf, i)
		_combine(buf, i+1, n, k, result)
		buf = buf[:len(buf)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
