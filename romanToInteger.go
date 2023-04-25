package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var result int
	len := len(s)
	for i := 0; i < len; i++ {
		switch s[i] {
		case 'I':
			if i+1 < len && (s[i+1] == 'V' || s[i+1] == 'X') {
				result -= 1
			} else {
				result += 1
			}
		case 'V':
			result += 5
		case 'X':
			if i+1 < len && (s[i+1] == 'L' || s[i+1] == 'C') {
				result -= 10
			} else {
				result += 10
			}
		case 'L':
			result += 50
		case 'C':
			if i+1 < len && (s[i+1] == 'D' || s[i+1] == 'M') {
				result -= 100
			} else {
				result += 100
			}
		case 'D':
			result += 500
		case 'M':
			result += 1000
		}
	}
	return result
}
