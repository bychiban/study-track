package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(stoneGameIII([]int{-1, -2, -3, -4}))
}

func stoneGameIII(piles []int) string {
	n := len(piles)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, 2)
		memo[i][1] = math.MinInt
	}
	memo[n-1][0] = piles[n-1]
	for i := len(piles) - 2; i >= 0; i-- {
		memo[i][0] = memo[i+1][0] + piles[i]
	}
	aliceScore := solve(&piles, 0, 3, &memo)
	bobScore := memo[0][0] - aliceScore
	if aliceScore > bobScore {
		return "Alice"
	} else if bobScore > aliceScore {
		return "Bob"
	}
	return "Tie"
}

func solve(piles *[]int, current int, m int, memo *[][]int) int {

	if current == len(*piles) {
		return 0
	}

	if (*memo)[current][1] != math.MinInt {
		return (*memo)[current][1]
	}

	result := math.MinInt
	for x := 1; x <= m && (current+x) <= len(*piles); x++ {
		result = max(result, (*memo)[current][0]-solve(piles, current+x, m, memo))
	}
	(*memo)[current][1] = result
	return result
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
