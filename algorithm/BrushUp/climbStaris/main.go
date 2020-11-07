package main

// 动态规划求解
// func climbStairs(n int) int {
// 	var a, b, c = 0, 0, 1
// 	for i := 1; i <= n; i++ {
// 		a = b
// 		b = c
// 		c = a + b
// 	}

// 	return c
// }

// 使用递归求解

//傻递归会超时,因为会生成很多的重复递归树
func climb2(n int) int {
	if n <= 2 {
		return n
	}
	return climb2(n-1) + climb2(n-2)
}

//递归+map作为缓存
func climbStairs(n int) int {

	result := make(map[int]int)

	var a func(n int, result map[int]int) int

	a = func(n int, result map[int]int) int {
		if n == 0 {
			return 1
		}
		if n < 0 {
			return 0
		}
		if result[n] > 0 {
			return result[n]
		}
		k := a(n-1, result) + a(n-2, result)
		result[n] = k
		return k
	}
	a(n, result)

	return result[n]
}

func main() {

}
