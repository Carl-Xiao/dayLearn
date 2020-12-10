package main

//判断有多少个1 消除低位的1
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count++
		num = num & (num - 1)
	}
	return count
}

func main() {

}
