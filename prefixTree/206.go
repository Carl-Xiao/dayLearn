package main

// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

// 示例:

// Trie trie = new Trie();

// trie.insert("apple");
// trie.search("apple");   // 返回 true
// trie.search("app");     // 返回 false
// trie.startsWith("app"); // 返回 true
// trie.insert("app");
// trie.search("app");     // 返回 true
// 说明:

// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。

//树节点
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = new(Trie)
		}
		this = this.next[v-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	if this.isEnd == false {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	return true
}
