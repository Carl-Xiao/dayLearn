package main

//17 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

//1 第一步将各个数字对应的字母缓存

//2 用递归 别去人脑递归,会把自己困在旋涡中
var combinations []string

func letterCombinations(digits string) []string {
	letters := make(map[string]string)
	letters["2"] = "abc"
	letters["3"] = "def"
	letters["4"] = "ghi"
	letters["5"] = "jkl"
	letters["6"] = "mno"
	letters["7"] = "pqrs"
	letters["8"] = "tuv"
	letters["9"] = "wxyz"

	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "", letters)
	return combinations
}

func backtrack(digits string, index int, combination string, letterMap map[string]string) {
	//跳出循环
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	}

	//处理逻辑
	digit := string(digits[index])
	letters := letterMap[digit]
	lettersCount := len(letters)
	for i := 0; i < lettersCount; i++ {
		backtrack(digits, index+1, combination+string(letters[i]), letterMap)
	}
}

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

//

// 示例 1:

// 输入: [3,2,3]
// 输出: 3
// 示例 2:

// 输入: [2,2,1,1,1,2,2]
// 输出: 2

func majorityElement(nums []int) int {
	var major, count = nums[0], 1
	for i := 1; i < len(nums); i++ {
		if major == nums[i] {
			count++
		} else if count == 0 {
			major = nums[i]
			count++
		} else {
			count--
		}
	}
	return major
}

func main() {
	majorityElement([]int{6, 5, 5})
}
