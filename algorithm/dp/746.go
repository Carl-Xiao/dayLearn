package main

import "fmt"

//使用最小花费爬楼梯

//题目完全不知道在讲什么东西
// 看了评论后才明白，  给一个数组
// 可以选择跳过一个楼梯 或者两个楼梯 求每层消耗最小的

//
//状态转移方程  dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	fmt.Println(dp)
	return dp[len(cost)]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{0, 0, 1, 1}))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
