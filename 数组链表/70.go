package main

import "fmt"

//爬楼梯

//递归题 f(n)=f(n-1)+f(n-2)
func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	var dp []int
	dp = append(dp, 1)
	dp = append(dp, 2)
	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	fmt.Println(dp)
	return dp[n-1]
}
func main() {
	fmt.Println(climbStairs(3))
}
