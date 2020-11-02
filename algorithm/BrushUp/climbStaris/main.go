package main

func climbStairs(n int) int {
	var a, b, c = 0, 0, 1
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}
func main() {

}
