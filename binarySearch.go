package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 12))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		m := (left + right) / 2
		if nums[m] < target {
			left = m + 1
		} else if nums[m] > target {
			right = m - 1
		} else {
			return m
		}
	}
	return -1
}
