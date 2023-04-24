package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome_2(11))
}

func isPalindrome_1(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome_2(x int) bool {
	if x < 0 {
		return false
	}
	right, left := 0, x
	for left > 0 {
		right = right*10 + left%10
		left = left / 10
	}
	return right == x
}
