package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{9, 10, 1, 7, 3}))
}

func lastStoneWeight(stones []int) int {
	i := len(stones) - 1
	sort.Ints(stones)
	for i > 0 {
		sub := stones[i] - stones[i-1]
		stones = stones[:i-1]
		i -= 1
		t := Search(stones, sub)
		stones = append(stones[:t+1], stones[t:]...)
		stones[t] = sub
	}

	return stones[0]
}

func Search(nums []int, target int) int {
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
