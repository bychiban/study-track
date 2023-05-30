package main

import (
	"fmt"
)

func main() {

	obj := Constructor()
	obj.Add(9)
	obj.Remove(19)
	param_1 := obj.Contains(1)
	fmt.Println(param_1)

}

type MyHashSet struct {
	capacity int
	bmp      *[]bool
	length   int
}

var primes []int = []int{
	3, 7, 11, 17, 23, 29, 37, 47, 59, 71, 89, 107, 131, 163, 197, 239, 293, 353, 431, 521, 631, 761, 919,
	1103, 1327, 1597, 1931, 2333, 2801, 3371, 4049, 4861, 5839, 7013, 8419, 10103, 12143, 14591,
	17519, 21023, 25229, 30293, 36353, 43627, 52361, 62851, 75431, 90523, 108631, 130363, 156437,
	187751, 225307, 270371, 324449, 389357, 467237, 560689, 672827, 807403, 968897, 1162687, 1395263,
	1674319, 2009191, 2411033, 2893249, 3471899, 4166287, 4999559, 5999471, 7199369}

func Constructor() MyHashSet {
	initCapacity := getPrime(0)
	initBmp := make([]bool, initCapacity)
	return MyHashSet{capacity: initCapacity, bmp: &initBmp}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	if this.length == this.capacity {
		this.resizeHashSet()
	}
	for key >= this.capacity {
		this.resizeHashSet()
	}
	(*this.bmp)[key%this.capacity] = true
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		(*this.bmp)[key%this.capacity] = false
	}
}

func (this *MyHashSet) Contains(key int) bool {
	hash := key % this.capacity
	if (*this.bmp)[hash] == true && key < this.capacity {
		return true
	}
	return false
}

func (this *MyHashSet) resizeHashSet() {
	newCapacity := getPrime(this.capacity << 1)
	newBmp := make([]bool, newCapacity)
	for i := 0; i < this.capacity; i++ {
		newBmp[i] = (*this.bmp)[i]
	}
	this.capacity = newCapacity
	this.bmp = &newBmp
}

func getPrime(min int) int {
	for i := 0; i < len(primes); i++ {
		prime := primes[i]
		if prime > min {
			return prime
		}
	}
	return -1
}
