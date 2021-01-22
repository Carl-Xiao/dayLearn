package main

// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

//

// 示例：

// s = "leetcode"
// 返回 0

// s = "loveleetcode"
// 返回 2

func firstUniqChar(s string) int {
	mac := make(map[byte]int)
	for _, v := range s {
		mac[byte(v)]++
	}
	//保证有序是第一个不重复的字符
	for i := range s {
		if mac[s[i]] == 1 {
			return i
		}
	}
	return -1
}
