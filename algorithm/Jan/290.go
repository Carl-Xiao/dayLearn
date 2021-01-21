package main

import "strings"

//双向连通性 用两个map处理数据
func wordPattern(pattern string, s string) bool {
	n := len(pattern)
	ss := strings.Split(s, " ")
	if n != len(ss) {
		return false
	}
	macPattern := make(map[byte]string)
	macWord := make(map[string]byte)

	for i, v := range pattern {
		ch := ss[i]
		if macWord[ch] > 0 && macWord[ch] != byte(v) || macPattern[byte(v)] != "" && macPattern[byte(v)] != ch {
			return false
		}
		macPattern[byte(v)] = ch
		macWord[ch] = byte(v)
	}
	return true
}
