package main

// 212. 单词搜索 II
// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。

// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

//前缀树+dfs解决问题

type triee struct {
	next [26]*triee
	word string
}

func construcTriee() *triee {
	return &triee{}
}
func (t *triee) inserWord(w string) {
	for _, v := range w {
		v = v - 'a'
		if t.next[v] == nil {
			t.next[v] = &triee{}
		}
		t = t.next[v]
	}
	t.word = w
}

func dfs(i, j int, board [][]byte, node *triee, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.next[c-'a'] == nil { //访问过了或者字典中没有
		return
	}
	node = node.next[c-'a']

	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}
	//递归调用
	board[i][j] = '#'
	dfs(i+1, j, board, node, result)
	dfs(i-1, j, board, node, result)
	dfs(i, j+1, board, node, result)
	dfs(i, j-1, board, node, result)
	board[i][j] = c
}

func findWords(board [][]byte, words []string) []string {
	t := construcTriee()
	//插入字典树
	for _, v := range words {
		t.inserWord(v)
	}
	result := make([]string, 0)
	//查询
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, t, &result)
		}
	}
	return result
}
