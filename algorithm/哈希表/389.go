package main

import "fmt"

// 给定两个字符串 s 和 t，它们只包含小写字母。

// 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

// 请找出在 t 中被添加的字母。
func findTheDifference(s string, t string) byte {
	mac := make(map[rune]int)
	for _, value := range s {
		mac[value]++
	}
	for _, value := range t {
		if c, ok := mac[value]; ok {
			tep := c - 1
			if tep == 0 {
				delete(mac, value)
			} else {
				mac[value] = tep
			}
		} else {
			return byte(value)
		}
	}
	for key := range mac {
		return byte(key)
	}
	return byte('a')
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
