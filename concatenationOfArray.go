package main

import (
	"fmt"
)

func main() {
	fmt.Println(getConcatenation([]int{1000, 2000, 4000, 3000}))
}

func getConcatenation(nums []int) []int {
	length := len(nums)
	result := make([]int, length*2)
	for i := 0; i < length; i++ {
		result[i], result[i+length] = nums[i], nums[i]
	}
	return result
}
