package main

import (
	"fmt"
)

func main() {
	fmt.Println(average([]int{1000, 2000, 4000, 3000}))
}

func average(salary []int) float64 {
	sum := 0
	max := salary[0]
	min := salary[0]
	for _, value := range salary {
		sum = sum + value
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return float64(sum-max-min) / float64(len(salary)-2)
}
