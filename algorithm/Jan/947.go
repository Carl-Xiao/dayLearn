package main

import "fmt"

// 如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头
func removeStones(stones [][]int) int {
	arr := make([]int, len(stones))
	size := make([]int, len(stones))
	for i := 0; i < len(stones); i++ {
		arr[i] = i
		size[i] = 1
	}
	s := &stone{
		arr:  arr,
		size: size,
	}
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				s.union(i, j)
			}
		}
	}
	//查看跟节点数量
	count := 0
	fmt.Println(s.arr)
	fmt.Println(s.size)
	for i, v := range s.arr {
		if i == v && s.size[i] > 1 {
			count += s.size[i] - 1
		}
	}
	return count
}

//并查集操作
type stone struct {
	arr  []int
	size []int
}

func (s *stone) find(p int) int {
	for p != s.arr[p] {
		s.arr[p] = s.arr[s.arr[p]]
		p = s.arr[p]
	}
	return p
}

func (s *stone) union(p, q int) {
	rootP := s.find(p)
	rootQ := s.find(q)
	if rootP == rootQ {
		return
	}
	s.arr[rootQ] = rootP
	s.size[rootP] += s.size[rootQ]
}
