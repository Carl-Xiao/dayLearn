package main

import "fmt"

//n最大取值
//我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//所有的丑数一定是2 3 5的乘积
func nthUglyNumber(n int) int {
	var a, b, c int

	var result int
	if n == 1 {
		return 1
	}
	for i := 1; i <= n; i++ {
		result = min(min(a*2, b*3), c*5)
		if result == a*2 {
			a++
		}
		if result == b*3 {
			b++
		}
		if result == c*5 {
			c++
		}
	}
	fmt.Println(result)
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nthUglyNumber(5)
}
