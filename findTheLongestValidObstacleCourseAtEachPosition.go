package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	result := make([]int, len(obstacles))
	subseq := []int{}

	for i := 0; i < len(obstacles); i++ {
		idx := searchInsert(subseq, obstacles[i])
		if idx == len(obstacles) {
			subseq = append(subseq, obstacles[i])
		} else {
			subseq = append(subseq[:idx], subseq[idx+1:len(subseq)]...)
			subseq[idx] = obstacles[i]
		}
		result[i] = idx + 1
	}

	return result
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var m int
	for left <= right {
		m = (left + right) / 2
		if nums[m] <= target {
			left = m + 1
		} else if nums[m] > target {
			right = m - 1
		}
	}

	return left
}

// comment
func longestObstacleCourseAtEachPositionStupid(obstacles []int) []int {
	result := make([]int, len(obstacles))

	for i := 0; i < len(obstacles); i++ {
		result[i] = 1
		for j := 0; j < i; j++ {
			if obstacles[i] >= obstacles[j] {
				result[i] = max(result[i], result[j]+1)
			}
		}
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
