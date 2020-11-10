package main

// 50. Pow(x, n)

// 实现 pow(x, n) ，即计算 x 的 n 次幂函数。

// 示例 1:

// 输入: 2.00000, 10
// 输出: 1024.00000
// 示例 2:

// 输入: 2.10000, 3
// 输出: 9.26100

//边界值的判断有点需要注意,其余没什么
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		return 1 / _pow(x, n)
	}
	return _pow(x, n)
}

func _pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := _pow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {

}
