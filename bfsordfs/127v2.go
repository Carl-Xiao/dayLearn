package main

import "fmt"

//双向dfs
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	dic := make(map[string]bool)
	for _, v := range wordList {
		dic[v] = true
	}
	if dic[endWord] == false {
		return 0
	}
	count := 0
	front := []string{beginWord}
	end := []string{endWord}

	for len(front) > 0 && len(end) > 0 {
		if len(front) > len(end) {
			front, end = end, front
		}
		temp := []string{}
		for _, word := range front {
			for j := 0; j < len(word); j++ {
				for k := 'a'; k <= 'z'; k++ {
					//替换j的元素
					tmp := word[:j] + string(k) + word[j+1:]

					if contains(tmp, end) {
						fmt.Println(tmp, end)
						return count + 1
					}

					if dic[tmp] == false {
						continue
					}
					delete(dic, tmp)
					temp = append(temp, tmp)
				}
			}
		}
		front = temp
		count++
	}
	return 0
}
func contains(key string, result []string) bool {
	for _, value := range result {
		if value == key {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(ladderLength2("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
