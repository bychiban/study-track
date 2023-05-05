package main

import (
	"fmt"
)

func main() {
	fmt.Println(canConstruct("aa", "ab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	hashMap := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		hashMap[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if value, ok := hashMap[ransomNote[i]]; ok && value > 0 {
			hashMap[ransomNote[i]]--
		} else {
			return false
		}
	}

	return true
}
