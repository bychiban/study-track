package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxScore([]int{39759, 619273, 859218, 228161, 944571, 597983, 483239, 179849, 868130, 909935, 912143, 817908, 738222, 653224}))
}

type Tuple struct {
	num1 int
	num2 int
}

func maxScore(nums []int) int {
	gcdHash := map[Tuple]int{}
	memo := map[string]int{}
	return solve(nums, 1, &gcdHash, &memo)
}

func solve(nums []int, multiplier int, gcdHash *map[Tuple]int, memo *map[string]int) int {
	max := 0

	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (*gcdHash)[Tuple{nums[i], nums[j]}] == 0 {
				(*gcdHash)[Tuple{nums[i], nums[j]}] = gcd(nums[i], nums[j])
			}
			newArray := remove(&nums, i, j)
			key := fmt.Sprintf("%d>%v", multiplier, newArray)
			if (*memo)[key] == 0 {
				(*memo)[key] = solve(
					newArray,
					multiplier+1,
					gcdHash,
					memo,
				)
			}
			current := (*gcdHash)[Tuple{nums[i], nums[j]}]*multiplier + (*memo)[key]

			if max < current {
				max = current
			}
		}
	}
	return max
}

func remove(nums *[]int, x, y int) []int {
	output := make([]int, len(*nums)-2)
	for i, j := 0, 0; j < len(*nums); j++ {
		if j != x && j != y {
			output[i] = (*nums)[j]
			i++
		}
	}
	return output
}

func gcd(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}
