package main

import "fmt"

func main() {
	fmt.Println(countGoodStrings(3, 3, 1, 1))
}

func countGoodStrings(low int, high int, zero int, one int) int {
	shift := zero + 1
	if shift < one {
		shift = one + 1
	}
	memo := make([]int, shift)
	memo[0] = 1
	result := 0
	for i := 0; i <= high; i++ {
		if i-zero >= 0 {
			memo[i%shift] = memo[(i-zero)%shift]
		}
		if i-one >= 0 {
			memo[i%shift] += memo[(i-one)%shift]
		}
		memo[i%shift] = mod(memo[i%shift])
		if i >= low {
			result = mod(result + memo[i%shift])
		}
	}

	return result
}

func mod(x int) int { return x % 1000000007 }
