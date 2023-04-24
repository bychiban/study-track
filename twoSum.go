package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, value := range nums {
		if j, ok := hashMap[target-value]; ok {
			return []int{i, j}
		}
		hashMap[value] = i
	}

	return []int{}
}
