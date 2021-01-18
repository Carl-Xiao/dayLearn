package main

//Trie 前缀树

//明白 什么是前缀树

//将对应的字符串用树的形式的存放,也不知道是谁想出来的这种结构

//知道就知道，不知道就写不出来这东西

//Trie 树
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */

//ConstructorTrie 构造函数
func ConstructorTrie() Trie {
	return Trie{}
}

//Insert 插入值
func (t *Trie) Insert(word string) {
	node := t
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			node.next[v] = &Trie{}
		}
		node = node.next[v]
	}
	node.isEnd = true
}

//Search 查询值
func (t *Trie) Search(word string) bool {
	node := t
	for _, v := range word {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return node.isEnd
}

// StartsWith 返回给定前缀是否存在
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, v := range prefix {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return true
}
