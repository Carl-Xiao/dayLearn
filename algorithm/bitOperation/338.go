package main

//dp方程
func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 1; i <= num; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}
