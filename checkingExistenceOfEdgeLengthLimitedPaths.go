package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(distanceLimitedPathsExist(3,
		[][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}},
		[][]int{{0, 1, 2}, {0, 2, 5}}))
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	uf := Constructor(n)
	m := len(queries)
	e := len(edgeList)
	answer := make([]bool, m)
	sorted_q := make([][]int, m)
	for i := 0; i < m; i++ {
		sorted_q[i] = []int{queries[i][0], queries[i][1], queries[i][2], i}
	}
	sort.Slice(sorted_q, func(i, j int) bool {
		return sorted_q[i][2] < sorted_q[j][2]
	})
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	edges_index := 0
	for _, group := range sorted_q {
		p := group[0]
		q := group[1]
		limit := group[2]
		idx := group[3]
		for edges_index < e && edgeList[edges_index][2] < limit {
			node1 := edgeList[edges_index][0]
			node2 := edgeList[edges_index][1]
			uf.union(node1, node2)
			edges_index += 1
		}
		answer[idx] = uf.connected(p, q)
	}

	return answer
}

type UnionFind struct {
	root []int
	rank []int
}

func Constructor(n int) UnionFind {
	root := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 1
	}
	return UnionFind{root, rank}
}

func (unionFind *UnionFind) find(x int) int {
	if x == unionFind.root[x] {
		return x
	}
	unionFind.root[x] = unionFind.find(unionFind.root[x])
	return unionFind.root[x]
}

func (unionFind *UnionFind) union(x, y int) {
	rootX := unionFind.find(x)
	rootY := unionFind.find(y)
	if rootX != rootY {
		if unionFind.rank[rootX] > unionFind.rank[rootY] {
			unionFind.root[rootY] = rootX
		} else if unionFind.rank[rootX] < unionFind.rank[rootY] {
			unionFind.root[rootX] = rootY
		} else {
			unionFind.root[rootY] = rootX
			unionFind.rank[rootX] += 1
		}
	}
}

func (unionFind *UnionFind) connected(x, y int) bool {
	return unionFind.find(x) == unionFind.find(y)
}
