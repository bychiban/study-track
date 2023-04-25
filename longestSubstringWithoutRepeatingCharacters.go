package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkewb"))
}

func lengthOfLongestSubstring(s string) int {
	var longestSubstring int
	hashMap := make(map[byte]int)
	length := len(s)
	var l int
	for r := 0; r < length; r++ {
		if value, ok := hashMap[s[r]]; ok && value >= l {
			longestSubstring = max(longestSubstring, r-l)
			l = value + 1
		}
		hashMap[s[r]] = r
	}
	return max(longestSubstring, length-l)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
