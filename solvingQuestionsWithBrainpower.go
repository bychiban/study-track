package main

import "fmt"

func main() {
	fmt.Println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
}

func mostPoints(questions [][]int) int64 {
	results := make([]int64, len(questions))

	for i := 0; i < len(questions); i++ {
		results[i] = -1
	}

	return solve(&questions, 0, &results)
}

func solve(questions *[][]int, n int, result *[]int64) int64 {
	if len(*questions) <= n {
		return 0
	}
	question := (*questions)[n]
	if (*result)[n] == -1 {
		(*result)[n] = max(int64(question[0])+solve(questions, n+question[1]+1, result), solve(questions, n+1, result))
	}

	return (*result)[n]
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
