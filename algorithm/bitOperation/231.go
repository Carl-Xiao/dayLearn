package main

// 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

func isPowerOfTwo(n int) bool {

	return n&(n-1) == 0

}
