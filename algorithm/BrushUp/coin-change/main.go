package main

import (
	"math"
	"sort"
)

// 322. 零钱兑换

//贪心算法

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

// 你可以认为每种硬币的数量是无限的。

//dfs+贪心的思想

//还是要递归才能解决问题，贪心更像是一种解题的思考

// 示例 1：

// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] > c[j]
}

func coinChange(coins []int, amount int) int {
	var a IntSlice
	a = coins
	sort.Sort(a)

	miniAns := math.MaxInt64

	__coinchange(a, amount, 0, 0, &miniAns)
	if miniAns == math.MaxInt64 {
		return -1
	}

	return miniAns
}

func __coinchange(coins []int, amount, index, count int, min *int) {
	if amount == 0 {
		if count < *min {
			*min = count
		}
		return
	}
	if index == len(coins) {
		return //遍历到最后,无解
	}
	for maxCoinNum := amount / coins[index]; maxCoinNum >= 0 && maxCoinNum+count < *min; maxCoinNum-- {
		__coinchange(coins, amount-maxCoinNum*coins[index], index+1, count+maxCoinNum, min)
	}
}

func main() {
	coinChange([]int{4, 1, 2}, 1)
}
