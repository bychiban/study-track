package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4, 5, 6}))
}

func runningSum(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	result[0] = nums[0]
	for i := 1; i < n; i++ {
		result[i] = result[i-1] + nums[i]
	}
	return result
}
