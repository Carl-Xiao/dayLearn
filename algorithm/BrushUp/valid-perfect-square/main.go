package main

import "fmt"

func isPerfectSquare(num int) bool {
	left := 1
	right := num / 2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		}
		if mid*mid > right {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
}
