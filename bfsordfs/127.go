package main

import "fmt"

//127 单词接龙

// 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

// 每次转换只能改变一个字母。
// 转换过程中的中间单词必须是字典中的单词。

// 输入:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]

// 输出: 5

// 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//      返回它的长度 5。
//单dfs会超时,其实也能理解因为从前到后要换26次 时间复杂度太高了
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool)
	for _, value := range wordList {
		wordMap[value] = true
	}
	count := 0
	queue := []string{}
	queue = append(queue, beginWord)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return count + 1
			}
			//单词变化
			wordBytes := []byte(word)
			for i := 0; i < len(wordBytes); i++ {
				for j := 'a'; j <= 'z'; j++ {
					c := wordBytes[i]
					if wordBytes[i] == byte(j) {
						continue
					}
					wordBytes[i] = byte(j)
					if isvald(string(wordBytes), wordMap) {
						queue = append(queue, string(wordBytes))
					}
					wordBytes[i] = c
				}

			}
			size--
		}
		count++

	}
	return 0
}

func isvald(word string, wordMap map[string]bool) bool {
	_, ok := wordMap[word]
	return ok
}

//双向dfs

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
