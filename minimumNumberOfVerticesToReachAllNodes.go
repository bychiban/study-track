package main

import "fmt"

func main() {
	fmt.Println(findSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	results := []int{}
	hashMap := make(map[int]bool)

	for _, item := range edges {
		hashMap[item[1]] = false
	}

	for i := 0; i < n; i++ {
		if _, ok := hashMap[i]; !ok {
			results = append(results, i)
		}
	}

	return results
}
