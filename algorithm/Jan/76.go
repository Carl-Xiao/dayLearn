package main

//利用滑动窗口解题
func minWindow(s string, t string) string {
	ori := map[byte]int{}
	for _, v := range t {
		ori[byte(v)]++
	}
	sLen := len(s)
	missingType := len(ori)

	l, r := 0, 0
	minLen := len(s) + 1
	start := len(s)

	for ; r < sLen; r++ {
		rightChar := s[r]
		if _, ok := ori[rightChar]; ok {
			ori[rightChar]--
			if ori[rightChar] == 0 {
				missingType--
			}
		}
		for missingType == 0 {
			if r-l+1 < minLen {
				minLen = r - l + 1
				start = l
			}
			leftChar := s[l]
			if _, ok := ori[leftChar]; ok {
				ori[leftChar]++
			}
			if ori[leftChar] > 0 {
				missingType++
			}
			l++
		}
	}
	//过滤边界值
	if start == len(s) {
		return ""
	}
	return s[start : start+minLen]
}
