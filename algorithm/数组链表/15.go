package main

//枚举完所有的例子 感觉不太好
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// func threeSum(nums []int) [][]int {
// 	var res [][]int
// 	m := make(map[int]int)
// 	m2 := make(map[string]bool)
// 	m3 := make(map[int]bool)
// 	for _, value := range nums {
// 		m3[value] = true
// 	}

// 	sort.Ints(nums)
// 	for i := 0; i < len(nums); i++ {
// 		if i > 0 && nums[i] == nums[i-1] { //三元组元素a去重
// 			continue
// 		}
// 		for j := i + 1; j < len(nums); j++ {
// 			if j > i+2 && nums[j] == nums[j-1] && nums[j-1] == nums[j-2] { // 三元组元素b去重
// 				continue
// 			}
// 			temp := 0 - (nums[i] + nums[j])
// 			if _, ok := m3[temp]; !ok {
// 				con
// 			}

// 			_, ok := m[temp]
// 			if ok {
// 				result := []int{nums[i], nums[j], temp}
// 				sort.Ints(result)
// 				key, _ := json.Marshal(result)
// 				if _, ok := m2[string(key)]; !ok {
// 					res = append(res, result)
// 					delete(m, temp)
// 					m2[string(key)] = true
// 				}
// 			} else {
// 				m[nums[j]] = nums[j]
// 			}
// 		}
// 	}
// 	return res
// }

// func main() {
// 	a := threeSum([]int{-1, 0, 1, 2, -1, -4})
// 	fmt.Println(a)
// }
