package main

import "fmt"

func addToArrayForm(A []int, K int) (ans []int) {
	for i := len(A) - 1; i >= 0; i-- {
		sum := A[i] + K%10
		K /= 10
		if sum >= 10 {
			K++
			sum -= 10
		}
		ans = append(ans, sum)
	}
	//k还有元素没有处理完
	for K > 0 {
		ans = append(ans, K%10)
		K /= 10
	}
	reverse(ans)
	return
}
func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-1-i], nums[i]
	}
}

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
}
