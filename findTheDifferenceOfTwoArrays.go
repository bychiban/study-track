package main

import "fmt"

func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{1, 2, 2, 7, 11, 16}
	fmt.Println(findDifference(nums1, nums2))
}

type IntSet map[int]bool

func findDifference(nums1 []int, nums2 []int) [][]int {
	intSet1 := make(IntSet)
	intSet2 := make(IntSet)
	for i := 0; i < len(nums1); i++ {
		intSet1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		id := nums2[i]
		if _, ok := intSet1[id]; ok {
			intSet1[id] = false
		} else {
			intSet2[id] = true
		}
	}

	return [][]int{intSet1.toSlice(), intSet2.toSlice()}
}

func (intSet *IntSet) toSlice() []int {
	keys := []int{}
	for k, value := range *intSet {
		if value == true {
			keys = append(keys, k)
		}
	}
	return keys
}
