package main

//颠倒给定的 32 位无符号整数的二进制位。

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	power := uint32(31)
	for num > 0 {
		result += (num & 1) << power
		num = num >> 1
		power--
	}
	return result
}
