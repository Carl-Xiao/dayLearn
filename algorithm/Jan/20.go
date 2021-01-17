package main

//利用栈求解
func isValid(s string) bool {
	var stack []byte
	eq := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for _, v := range []byte(s) {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				sv := stack[len(stack)-1]
				if eq[sv] != v {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
