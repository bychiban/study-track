package main

import (
	"fmt"
)

func main() {
	fmt.Println(numTrees(19))
}

func numTrees(n int) int {
	var hashMap map[int]int = make(map[int]int, 0)
	hashMap[0] = 1
	hashMap[1] = 1
	return numTreesCalc(n, &hashMap)
}

func numTreesCalc(n int, results *map[int]int) int {
	if value, ok := (*results)[n]; ok {
		return value
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + numTreesCalc(i-1, results)*numTreesCalc(n-i, results)
	}
	(*results)[n] = sum
	return sum
}
