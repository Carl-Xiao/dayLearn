package main

//爬楼梯

//这个题感觉用动态规划解题更合理一些

//用递归 f(n)=f(n-1)+f(n-2)
func climbStairs(n int) int {

	//边界值尽量的完善
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
