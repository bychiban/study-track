package main

import (
	"math"
)

func main() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	kthLargest.Add(3)  // return 4
	kthLargest.Add(5)  // return 5
	kthLargest.Add(10) // return 5
	kthLargest.Add(9)  // return 8
	kthLargest.Add(4)  // return 8
}

type KthLargest struct {
	sortedArray []int
	len         int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{sortedArray: make([]int, k), len: k}
	for i := 0; i < k; i++ {
		kthLargest.sortedArray[i] = math.MinInt
	}
	for i := 0; i < len(nums); i++ {
		kthLargest.Add(nums[i])
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	left := 0
	right := this.len - 1
	for left <= right {
		mid := left + (right-left)/2
		if this.sortedArray[mid] == val {
			this.sortedArray = append(this.sortedArray[1:mid+1], this.sortedArray[mid:]...)
			return this.sortedArray[0]
		}
		if this.sortedArray[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == 0 {
		return this.sortedArray[0]
	}
	if left == this.len {
		copy(this.sortedArray, this.sortedArray[1:left])
		this.sortedArray[left-1] = val
		return this.sortedArray[0]
	}
	copy(this.sortedArray, this.sortedArray[1:left])
	this.sortedArray[left-1] = val
	return this.sortedArray[0]
}
