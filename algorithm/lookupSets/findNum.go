package main

// 2. 老乡群
// 某公司有 N 名员工，其中有些是来自同一个地方的，有些则不是。他们的老乡关系具有传递性。

// 如果已知 A 是 B 是老乡，B 是 C 也是老乡，那么我们可以认为 A 跟 C 也是老乡。所谓的老乡群，就是指来自同一个地方的员工集合。

// 给定一个 N * N 的矩阵 M，表示员工之间的老乡关系。
// 如果 Mi = 1，表示已知第 i 个和第 j 个是老乡，否则为不知道。
// 请输出所有员工中的已知的老乡群总数。

// 示例

// 输入：
// [[1,1,0],
// [1,1,0],
// [0,0,1]]

// 输出：2

// 解释：
// 已知员工 0 和员工 1 互为老乡，他们在一个老乡群。
// 员工 2 自己在一个老乡群。所以返回 2 。

//求合并的东西~用并查集求解
type NumUniton struct {
	count  int
	parent []int
}

func constructorNum(n int) *NumUniton {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return &NumUniton{
		count:  n,
		parent: nums,
	}
}

//找到最终的父节点
func (u *NumUniton) find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (num *NumUniton) union(p, q int) {
	rootP := num.find(p)
	rootQ := num.find(q)

	if rootP == rootQ {
		return
	}
	num.parent[rootP] = rootQ
	num.count--
}

func findNum(M [][]int) int {
	m := len(M)
	n := len(M[0])

	friend := constructorNum(m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				friend.union(i, j)
			}
		}
	}
	return friend.count
}
