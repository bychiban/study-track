package main

import (
	"fmt"
)

func main() {
	fmt.Println(diagonalSum([][]int{{3, 1, 5}, {6, 4, 2}, {1, 1, 1}}))
}

func diagonalSum(mat [][]int) int {
	sum := 0
	n := len(mat)
	for i := 0; i < n; i++ {
		sum = sum + mat[i][i] + mat[i][n-1-i]
	}
	if n%2 != 0 {
		sum = sum - mat[n/2][n/2]
	}
	return sum
}
