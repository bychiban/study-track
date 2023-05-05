package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
}

func maximumWealth(accounts [][]int) int {
	maxSum := 0
	for _, account := range accounts {
		sum := 0
		for _, value := range account {
			sum += value
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
