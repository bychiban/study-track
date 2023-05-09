package main

import (
	"fmt"
)

func main() {
	fmt.Println(spiralOrder([][]int{
		{3, 1, 5, 7},
		{6, 4, 2, 5},
		{1, 1, 1, 4}}))
}

func spiralOrder(matrix [][]int) []int {
	xn := len(matrix[0])
	yn := len(matrix)
	result := make([]int, xn*yn)

	goLeft := func(matrix [][]int, x0, y0, x1, y1 int, result *[]int, i *int) {
		for ; x0 <= x1; x0++ {
			(*result)[*i] = matrix[y0][x0]
			*i++
		}
	}
	goDown := func(matrix [][]int, x0, y0, x1, y1 int, result *[]int, i *int) {
		for ; y0 <= y1; y0++ {
			(*result)[*i] = matrix[y0][x1]
			*i++
		}
	}
	goRight := func(matrix [][]int, x0, y0, x1, y1 int, result *[]int, i *int) {
		for ; x1 >= x0; x1-- {
			(*result)[*i] = matrix[y1][x1]
			*i++
		}
	}
	goUp := func(matrix [][]int, x0, y0, x1, y1 int, result *[]int, i *int) {
		for ; y1 >= y0; y1-- {
			(*result)[*i] = matrix[y1][x0]
			*i++
		}
	}
	i, x0, y0, x1, y1 := 0, 0, 0, xn-1, yn-1

	for {
		goLeft(matrix, x0, y0, x1, y1, &result, &i)
		y0++
		if i == len(result) {
			return result
		}
		goDown(matrix, x0, y0, x1, y1, &result, &i)
		x1--
		if i == len(result) {
			return result
		}
		goRight(matrix, x0, y0, x1, y1, &result, &i)
		y1--
		if i == len(result) {
			return result
		}
		goUp(matrix, x0, y0, x1, y1, &result, &i)
		x0++
		if i == len(result) {
			return result
		}
	}
}
