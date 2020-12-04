package main

//前缀树 先构造一个前缀树再处理数据

type TrieNode struct {
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	//先构造前缀树
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}
	result := make([]string, 0)
	//使用dfs遍历二维数组
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, board, root, &result)
		}
	}
	return result
}

func dfs(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	//剪枝
	if c == '#' || node.children[c-'a'] == nil {
		return
	}
	node = node.children[c-'a']
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = "" // 防止重复添加
	}
	board[i][j] = '#'
	//右
	dfs(i+1, j, board, node, result)
	//左
	dfs(i-1, j, board, node, result)
	//下
	dfs(i, j+1, board, node, result)
	//上
	dfs(i, j-1, board, node, result)
	board[i][j] = c
}
