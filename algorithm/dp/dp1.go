package main

import "fmt"

// 小张和骑友小王在路上上不小心走散，小张通过定位仪找到小王的位置并且希望能快速找到小王。
// 小张最开始在 N 点位置 (0 ≤ N ≤ 100,000)，小王显示在 K 点位置 (0 ≤ K ≤ 100,000)，
// 小张有两种移动方式：
// 方式一：在任意点 X，向前走一步 X + 1 或向后走一步 X - 1 需要花费一分钟
// 方式二：在任意点 X，向前移动到 2 * X 位置需要花费一分钟

func findMinMinutes(n int, k int) int {
	distance := 0
	if n < k {
		distance = k - n
	} else {
		distance = n - k
		return distance
	}
	steps := 1
	sum := 0
	for steps = 1; steps < distance; steps = steps * 2 {
		sum = sum + 1
	}
	return sum
}

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

func findNum(M [][]int) int {
	count := 0
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if i == 0 || j == 0 {
				if M[i][j] != M[i][j+1] && M[i][j] != M[i+1][j] {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	fmt.Println(findMinMinutes(4, 7))
}
