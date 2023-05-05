package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

func maxVowels(s string, k int) int {
	left := 0
	right := k - 1
	currentValue := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currentValue++
		}
	}
	maxValue := currentValue
	for right = right + 1; right < len(s); left, right = left+1, right+1 {
		if isVowel(s[left]) {
			currentValue--
		}
		if isVowel(s[right]) {
			currentValue++
		}
		if currentValue > maxValue {
			maxValue = currentValue
			if maxValue == k {
				return maxValue
			}
		}
	}
	return maxValue
}

func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
