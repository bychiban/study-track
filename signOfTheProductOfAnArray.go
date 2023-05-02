package main

import (
	"fmt"
)

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, -1}))
}

func arraySign(nums []int) int {
	result := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			result = -result
		}
	}
	return result
}
