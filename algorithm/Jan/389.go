package main

//389找不同

//将两个字符串进行排序,然后逐一对比

//使用map处理
func findTheDifference(s string, t string) byte {
	var a byte
	mac := make(map[byte]int)
	for _, v := range t {
		mac[byte(v)]++
	}

	for _, v := range s {
		mac[byte(v)]--
	}

	for i, v := range mac {
		if v > 0 {
			return i
		}
	}
	return a
}
