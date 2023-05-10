package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(1))
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	goLeft := func(x0, y0, x1, y1 int, matrix *[][]int, i *int) {
		for ; x0 <= x1; x0++ {
			(*matrix)[y0][x0] = *i
			*i++
		}
	}
	goDown := func(x0, y0, x1, y1 int, matrix *[][]int, i *int) {
		for ; y0 <= y1; y0++ {
			(*matrix)[y0][x1] = *i
			*i++
		}
	}
	goRight := func(x0, y0, x1, y1 int, matrix *[][]int, i *int) {
		for ; x1 >= x0; x1-- {
			(*matrix)[y1][x1] = *i
			*i++
		}
	}
	goUp := func(x0, y0, x1, y1 int, matrix *[][]int, i *int) {
		for ; y1 >= y0; y1-- {
			(*matrix)[y1][x0] = *i
			*i++
		}
	}
	x0, y0, x1, y1 := 0, 0, n-1, n-1
	i := 1
	for {
		goLeft(x0, y0, x1, y1, &matrix, &i)
		y0++
		if i-1 == n*n {
			return matrix
		}
		goDown(x0, y0, x1, y1, &matrix, &i)
		x1--
		if i-1 == n*n {
			return matrix
		}
		goRight(x0, y0, x1, y1, &matrix, &i)
		y1--
		if i-1 == n*n {
			return matrix
		}
		goUp(x0, y0, x1, y1, &matrix, &i)
		x0++
		if i-1 == n*n {
			return matrix
		}
	}
}
