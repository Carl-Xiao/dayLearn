package main

import "fmt"

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := new([]string)
	_letterCombinations("", 0, digits, result)
	return *result
}

var (
	letterMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
)

func _letterCombinations(buf string, index int, digits string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, buf)
		return
	}
	temp := string(digits[index])
	tempString := letterMap[temp]
	for _, value := range tempString {
		_letterCombinations(buf+string(value), index+1, digits, result)
	}

}

func main() {
	fmt.Println(letterCombinations("23"))
}
