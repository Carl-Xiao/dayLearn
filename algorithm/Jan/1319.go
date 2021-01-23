package main

//这么一看就算求并查集的联通次数

type connecte struct {
	parent []int
	size   int
}

func Constructorconnect(n int) *connecte {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &connecte{
		parent: parent,
		size:   n,
	}
}

func (c *connecte) find(p int) int {
	for p != c.parent[p] {
		c.parent[p] = c.parent[c.parent[p]]
		p = c.parent[p]
	}
	return p
}

func (c *connecte) union(p, q int) {
	rootP := c.find(p)
	rootQ := c.find(q)
	if rootP == rootQ {
		return
	}
	c.parent[rootP] = rootQ
	c.size--
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	c := Constructorconnect(n)
	for i := 0; i < len(connections); i++ {
		c.union(connections[i][0], connections[i][1])
	}
	return c.size - 1
}
