package main

//想到用栈来解决问题

//调试的时候需要多加边界值才能确保没有问题出现
func isValid(s string) bool {
	var stack []byte
	result := []byte(s)
	for i := 0; i < len(result); i++ {
		if result[i] == '(' {
			stack = append(stack, ')')
		} else if result[i] == '[' {
			stack = append(stack, ']')
		} else if result[i] == '{' {
			stack = append(stack, '}')
		} else {
			if len(stack) > 0 {
				temp := stack[len(stack)-1]
				if temp != result[i] {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
