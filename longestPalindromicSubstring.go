package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	palindromeLength := 0
	resultLeft := 0
	resultRight := 0
	left := 0
	right := 0
	for i := 0; i < (len(s) - palindromeLength/2); i++ {
		for left, right = i, i; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] != s[right] {
				break
			}
		}
		if palindromeLength < right-left-1 {
			palindromeLength = right - left - 1
			resultLeft = left + 1
			resultRight = right - 1
		}
		for left, right = i, i+1; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] != s[right] {
				break
			}
		}
		if palindromeLength < right-left-1 {
			palindromeLength = right - left - 1
			resultLeft = left + 1
			resultRight = right - 1
		}

	}

	return s[resultLeft : resultRight+1]
}
