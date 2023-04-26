package main

import (
	"fmt"
)

func main() {
	fmt.Println(addDigits(18))
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}

	return num % 9
}
