package main

//前缀树的实际应用
//请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

// 1 <= word.length <= 500
// addWord 中的 word 由小写英文字母组成
// search 中的 word 由 '.' 或小写英文字母组成 word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。

type WordDictionary struct {
	next  [26]*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for _, v := range word {
		v = v - 'a'
		if this.next[v] == nil {
			this.next[v] = &WordDictionary{}
		}
		this = this.next[v]
	}
	this.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var searchHelp func(word string, node *WordDictionary) bool
	searchHelp = func(word string, node *WordDictionary) bool {
		for i, v := range word {
			if v == '.' {
				//递归处理问题
				for _, dictionary := range node.next {
					if dictionary != nil {
						if searchHelp(word[i+1:], dictionary) { //遇到'.'的时候跳过'.'继续匹配接下来的字符，树的深度进1
							return true
						}
					}
				}
				return false
			}
			if node = node.next[v-'a']; node == nil {
				return false
			}
		}
		return node.isEnd
	}
	return searchHelp(word, this)
}
