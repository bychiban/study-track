package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 < len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	lo, hi := 0, len2*2
	for lo <= hi {
		mid2 := (lo + hi) / 2
		mid1 := len1 + len2 - mid2

		var L1 float64
		if mid1 == 0 {
			L1 = math.MinInt
		} else {
			L1 = float64(nums1[(mid1-1)/2])
		}
		var L2 float64
		if mid2 == 0 {
			L2 = math.MinInt
		} else {
			L2 = float64(nums2[(mid2-1)/2])
		}
		var R1 float64
		if mid1 == len1*2 {
			R1 = math.MaxInt
		} else {
			R1 = float64(nums1[(mid1)/2])
		}
		var R2 float64
		if mid2 == len2*2 {
			R2 = math.MaxInt
		} else {
			R2 = float64(nums2[(mid2)/2])
		}

		if L1 > R2 {
			lo = mid2 + 1
		} else if L2 > R1 {
			hi = mid2 - 1
		} else {
			return (math.Max(L1, L2) + math.Min(R1, R2)) / 2
		}
	}
	return -1
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	current := 0
	previous := 0
	for idx, i, j := 0, 0, 0; idx <= ((len1 + len2) / 2); idx++ {
		previous = current
		switch {
		case i == len1:
			current = nums2[j]
			j++
		case j == len2:
			current = nums1[i]
			i++
		case nums1[i] > nums2[j]:
			current = nums2[j]
			j++
		default:
			current = nums1[i]
			i++
		}
	}
	if (len1+len2)%2 == 0 {
		return float64(previous+current) / 2
	} else {
		return float64(current)
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	combined := []int{}
	var init *[]int
	var insert *[]int
	if len(nums1) > len(nums2) {
		init, insert = &nums1, &nums2
	} else {
		init, insert = &nums2, &nums1
	}

	lenInit := len(*init)
	lenInsert := len(*insert)

	for i := 0; i < lenInit; i++ {
		combined = append(combined, (*init)[i])
	}

	currentLength := lenInit

	for i := 0; i < lenInsert; i++ {
		target := (*insert)[i]
		idx := 0
		l, r := 0, currentLength-1
		for l <= r {
			idx = l + (r-l)/2
			if combined[idx] > target {
				r = idx - 1
			} else {
				l = idx + 1
			}
			if combined[idx] == target {
				break
			}
		}
		if l == currentLength {
			combined = append(combined, target)
		} else {
			combined = append(combined[:l+1], combined[l:currentLength]...)
			combined[l] = target
		}
		currentLength = currentLength + 1
	}
	idx := (lenInit + lenInsert) / 2
	if (lenInit+lenInsert)%2 == 0 {
		return float64(combined[idx-1]+combined[idx]) / 2
	} else {
		return float64(combined[idx])
	}
}
