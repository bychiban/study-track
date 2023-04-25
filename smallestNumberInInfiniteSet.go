package main

import "fmt"

func main() {
	obj := Constructor()
	param_1 := obj.PopSmallest()
	obj.AddBack(3)
	obj.AddBack(3)
	obj.AddBack(1)
	obj.AddBack(3)
	fmt.Println(param_1)
}

type SmallestInfiniteSet struct {
	excludedSet map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{excludedSet: make(map[int]bool)}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; ; i++ {
		if isExcluded, _ := this.excludedSet[i]; !isExcluded {
			this.excludedSet[i] = true
			return i
		}
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if isExcluded, ok := this.excludedSet[num]; ok && isExcluded {
		delete(this.excludedSet, num)
	}
}
