package main

// 剑指 Offer 49. 丑数

// 示例:

// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

func nthUglyNumber(n int) int {
	result := make([]int, n, n)
	result[0] = 1
	var j, k, l int
	for i := 1; i < n; i++ {
		a := result[j] * 2
		b := result[k] * 3
		c := result[l] * 5
		result[i] = min(min(a, b), c)
		if result[i] == a {
			j++
		}
		if result[i] == b {
			k++
		}
		if result[i] == c {
			l++
		}
	}
	return result[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
