package main

// 使用递归求解

// > 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// >
// > 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// >
// > 注意：给定 n 是一个正整数。
// >
// > 示例 1：
// >
// > 输入： 2
// > 输出： 2
// > 解释： 有两种方法可以爬到楼顶。
// > 1.  1 阶 + 1 阶
// > 2.  2 阶
// > 示例 2：

func climbStairs(n int) int {
	resultMap := make(map[int]int)

	_climbStairs(n, resultMap)

	return resultMap[n]
}

func _climbStairs(n int, resultMap map[int]int) int {
	//边界情况
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if resultMap[n] > 0 {
		return resultMap[n]
	}
	//处理逻辑
	k := _climbStairs(n-1, resultMap) + _climbStairs(n-2, resultMap)
	resultMap[n] = k
	return k
}

func main() {

}
