package main

import (
	"fmt"
)

func main() {
	fmt.Println(bulbSwitch(0))
	fmt.Println(bulbSwitch(1))
	fmt.Println(bulbSwitch(2))
	fmt.Println(bulbSwitch(99999))
}

func bulbSwitch(n int) int {
	result := 0
	for i := 1; i*i <= n; i++ {
		result++
	}
	return result
}
