package main

//用hash结构存储
func isAnagram(s string, t string) bool {
	mac := make(map[rune]int)

	for _, value := range s {
		mac[value]++
	}

	for _, value := range t {
		if c, ok := mac[value]; ok {
			c = c - 1
			if c == 0 {
				delete(mac, value)
			} else {
				mac[value] = c
			}
		} else {
			return false
		}
	}

	return len(mac) == 0

}
