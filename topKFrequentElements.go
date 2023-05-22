package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

type Tuple struct {
	count, num int
}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		hashMap[nums[i]]++
	}
	sortedArray := []Tuple{}
	for num, count := range hashMap {
		idx := searchInsert(&sortedArray, count)
		if idx >= len(sortedArray) {
			sortedArray = append(sortedArray, Tuple{count, num})
		} else {
			sortedArray = append(sortedArray[:idx+1], sortedArray[idx:]...)
			sortedArray[idx] = Tuple{count, num}
		}
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = sortedArray[len(sortedArray)-1-i].num
	}
	return result
}

func searchInsert(array *[]Tuple, target int) int {
	left := 0
	right := len(*array) - 1
	for left <= right {
		m := (left + right) / 2
		if (*array)[m].count < target {
			left = m + 1
		} else if (*array)[m].count > target {
			right = m - 1
		} else {
			return m
		}
	}
	return left
}
