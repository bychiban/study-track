package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchInsert([]int{}, 0))
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var m int
	for left <= right {
		m = (left + right) / 2
		if nums[m] < target {
			left = m + 1
		} else if nums[m] > target {
			right = m - 1
		} else {
			return m
		}
	}

	return left
}
