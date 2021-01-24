package main

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//主要掌握滑动窗口
//遍历整串
//左边
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	l, r, ans := 0, 0, 0
	for l < n {
		if r < n && m[s[r]] == 0 {
			m[s[r]] = 1
			r++
		} else {
			m[s[l]] = 0
			l++
		}
		ans = max(ans, r-l)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
