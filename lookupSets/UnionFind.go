package main

//https://leetcode-cn.com/problems/friend-circles/solution/union-find-suan-fa-xiang-jie-by-labuladong/

//并查集模板

type UnionFind struct {
	count  int
	parent []int
}

func constructor(n int) *UnionFind {
	count := n
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		count:  count,
		parent: parent,
	}

}

func (u *UnionFind) find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *UnionFind) uniton(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootP == rootQ {
		return
	}
	u.parent[rootP] = rootQ
	u.count--
}

func findCircleNum(M [][]int) int {
	n := len(M)

	union := constructor(n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				union.uniton(i, j)
			}
		}
	}
	return union.count
}
