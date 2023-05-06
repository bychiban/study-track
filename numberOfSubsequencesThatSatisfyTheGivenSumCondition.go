package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numSubseq([]int{27, 21, 14, 2, 15, 1, 19, 8, 12, 24, 21, 8, 12, 10, 11, 30, 15, 18, 28, 14, 26,
		9, 2, 24, 23, 11, 7, 12, 9, 17, 30, 9, 28, 2, 14, 22, 19, 19, 27, 6, 15, 12, 29, 2, 30, 11, 20, 30, 21,
		20, 2, 22, 6, 14, 13, 19, 21, 10, 18, 30, 2, 20, 28, 22}, 31))
}

func numSubseq(nums []int, target int) int {
	sort.Ints(nums[:])
	left := 0
	right := 0
	result := 0
	for left = 0; left < len(nums) && nums[left]<<1 <= target; left++ {
		current := nums[left]
		right = searchInsert(nums, target-nums[left])
		if right < left {
			return result % 1000000007
		} else if right > left {
			result += powModulo(right - left)
		} else if right == left && current<<1 <= target {
			result++
		}
	}

	return result % 1000000007
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

	if left > 0 {
		return left - 1
	}
	return left
}

func powModulo(n int) int {
	mod := 1000000007
	product := 1
	temp := 2
	for n != 0 {
		if n&1 == 1 {
			product = product * temp % mod
		}
		temp = (temp * temp) % mod
		n = n >> 1
	}
	return product
}
