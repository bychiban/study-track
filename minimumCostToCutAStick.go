package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minCost(9, []int{5, 6, 1, 4, 2}))
}

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	length := len(cuts)
	memo := make([][]int, length+1)
	for i := 0; i < length+1; i++ {
		memo[i] = make([]int, length+1)
		for j := 0; j < length+1; j++ {
			memo[i][j] = -1
		}
	}
	return solve(n, cuts, 0, 0, length-1, &memo)
}

func solve(n int, cuts []int, init, left, right int, memo *[][]int) int {
	length := len(cuts)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return n
	}
	if (*memo)[left][right] != -1 {
		return (*memo)[left][right]
	}
	result := math.MaxInt
	for i := 0; i < length; i++ {
		cut := cuts[i]
		result = min(
			result,
			n+
				solve(cut-init, cuts[0:i], init, left, left+i-1, memo)+
				solve(n+init-cut, cuts[i+1:], cut, left+i+1, right, memo))
	}
	(*memo)[left][right] = result
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
