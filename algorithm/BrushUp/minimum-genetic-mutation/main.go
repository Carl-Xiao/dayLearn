package main

import "fmt"

//433. 最小基因变化

//其中每个字符都属于 "A", "C", "G", "T"中的任意一个。  从start 变化为end，需要多少次

//变化的过程必须是属于bank

//如何思考

//第一遍 日常读不懂题目,这tm什么意思啊
//第二遍 看不懂题解
//第三遍 bfs广度优先遍历？？
//第四遍 从哪儿看出来的？
//第五遍 硬着头皮写
func minMutation(start string, end string, bank []string) int {
	chars := []byte{'A', 'C', 'G', 'T'}
	bankMap := make(map[string]bool, len(bank))
	visitedMap := make(map[string]bool, len(bank))
	for _, value := range bank {
		bankMap[value] = true
	}
	level := 0

	stack := []string{}
	stack = append(stack, start)
	for len(stack) != 0 {
		size := len(stack)

		for i := 0; i < size; i++ {
			fmt.Println(stack)
			element := stack[0]
			stack = stack[1:]
			fmt.Println(stack)

			if element == end {
				return level
			}
			tempByte := []byte(element)
			for k := range tempByte {
				old := tempByte[k]
				for _, v := range chars {
					tempByte[k] = v
					next := string(tempByte)
					if _, ok := bankMap[next]; ok {
						if _, ok := visitedMap[next]; !ok {
							visitedMap[next] = true
							stack = append(stack, next)
						}
					}
				}
				tempByte[k] = old
			}
		}
		level++
	}
	return -1
}

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
}
