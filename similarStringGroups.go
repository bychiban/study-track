package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSimilarGroups([]string{"rats", "rats", "arts", "star"}))
}

func numSimilarGroups(strs []string) int {
	length := len(strs)
	parents := make([]int, length)
	for i := 0; i < length; i++ {
		parents[i] = i
	}
	ds := DisjointSet{Parents: parents, Size: length}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if isSimilar(strs[i], strs[j]) {
				ds.Union(i, j)
			}
		}
	}

	return ds.Size
}

type DisjointSet struct {
	Parents []int
	Size    int
}

func (this *DisjointSet) Find(x int) int {
	if x == this.Parents[x] {
		return x
	}
	return this.Find(this.Parents[x])
}

func (this *DisjointSet) Union(x int, y int) int {
	rx, ry := this.Find(x), this.Find(y)
	if rx != ry {
		this.Parents[ry] = rx
		this.Size--
	}
	return rx
}

func isSimilar(x, y string) bool {
	missmatch := 0
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			missmatch++
			if missmatch > 2 {
				return false
			}
		}
	}
	return true
}
