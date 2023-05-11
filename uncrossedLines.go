package main

import "fmt"

func main() {
	fmt.Println(maxUncrossedLines([]int{3, 3}, []int{1, 3, 2, 1}))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 < n2 {
		return maxUncrossedLines(nums2, nums1)
	}

	prev, results := make([]int, n2+1), make([]int, n2+1)
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if nums1[i-1] == nums2[j-1] {
				results[j] = 1 + prev[j-1]
			} else {
				results[j] = max(results[j-1], prev[j])
			}
		}

		copy(prev, results)
	}
	return results[n2]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
