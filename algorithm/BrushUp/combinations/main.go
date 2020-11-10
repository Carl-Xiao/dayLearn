package main

import "fmt"

//NO.77 题目
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

//不要去人脑递归
//找共同点

func combine(n int, k int) [][]int {
	result := &[][]int{}
	_combine([]int{}, 1, n, k, result)
	return *result
}

func _combine(buf []int, begin, n, k int, result *[][]int) {
	if len(buf) == k {
		temp := make([]int, k)
		copy(temp, buf)
		*result = append(*result, temp)
		return
	}
	//处理逻辑
	for i := begin; i <= n; i++ {
		//修改状态
		buf = append(buf, i)
		_combine(buf, i+1, n, k, result)
		//还原状态
		buf = buf[:len(buf)-1]
	}

}

func main() {
	fmt.Println(combine(4, 2))
}
