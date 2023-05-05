package main

import "fmt"

func main() {
	fmt.Println(convert("AB", 1))
}

func convert(s string, numRows int) string {
	arr := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = []byte{}
	}
	x := 0
	y := 0

	goDown := func(char byte) {
		arr[y] = append(arr[y], char)
		if y < numRows-1 {
			y = y + 1
		}
	}
	goDiag := func(char byte) {
		arr[y] = append(arr[y], char)
		x = x + 1
		y = y - 1
	}
	cont := goDown
	for i := 0; i < len(s); i++ {
		if y == 0 {
			goDown(s[i])
			cont = goDown
		} else if y >= numRows-1 {
			goDiag(s[i])
			cont = goDiag
		} else {
			cont(s[i])
		}
	}
	var result string
	for i := 0; i < len(arr); i++ {
		result = result + string(arr[i])
	}
	return result
}
