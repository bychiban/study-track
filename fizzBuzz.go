package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		switch {
		case (i+1)%15 == 0:
			result[i] = "FizzBuzz"
		case (i+1)%3 == 0:
			result[i] = "Fizz"
		case (i+1)%5 == 0:
			result[i] = "Buzz"
		default:
			result[i] = strconv.Itoa(i + 1)
		}
	}
	return result
}
