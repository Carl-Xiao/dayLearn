package main

//递归的问题，就是不能去想太深，太深脑子就回不来了
func generateParenthesis(n int) []string {
	result := new([]string)

	_generateParenthesis("", 2*n, n, n, result)

	return *result
}

func _generateParenthesis(buf string, length, left, right int, result *[]string) {
	//移除条件,当长度等于2*n的时候，代表当前已经添加完毕
	if len(buf) == length {
		*result = append(*result, buf)
		return
	}
	//任意时候都可以
	if left > 0 {
		_generateParenthesis(buf+"(", length, left-1, right, result)
	}

	if left < right {
		_generateParenthesis(buf+")", length, left, right-1, result)
	}

}
