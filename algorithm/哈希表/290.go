package main

import "strings"

// 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
//提示: 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	word2ch := make(map[string]byte)
	ch2word := make(map[byte]string)

	//如果长度不一样~规律也对应不上
	if len(pattern) != len(words) {
		return false
	}

	for i, value := range words {
		ch := pattern[i]
		if word2ch[value] > 0 && ch2word[ch] != value || ch2word[ch] != "" && ch2word[ch] != value {
			return false
		}
		word2ch[value] = ch
		ch2word[ch] = value
	}
	return true
}


