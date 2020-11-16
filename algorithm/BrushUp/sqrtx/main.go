package main

//69. x 的平方根
func mySqrt(x int) int {
	left := 1
	right := x
	var result int
	for left <= right {
		mid := left + (right-1)/2
		if mid*mid <= x {
			left = left + 1
			result = mid
		} else {
			right = right - 1
		}

	}
	return result
}

// //牛顿迭代法
// func mySqrt(x int) int {
// 	if x == 0 {
// 		return 0
// 	}
// 	C, x0 := float64(x), float64(x)
// 	for {
// 		xi := 0.5 * (x0 + C/x0)
// 		if math.Abs(x0-xi) < 1e-7 {
// 			break
// 		}
// 		x0 = xi
// 	}
// 	return int(x0)
// }

func main() {

}
