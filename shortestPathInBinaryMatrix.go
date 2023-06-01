package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	length := len(grid)
	if grid[0][0] == 1 {
		return -1
	}
	listOfMoves := [8][2]int{{1, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][2]int{{0, 0}}
	grid[0][0] = 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currX, currY := current[0], current[1]
		if currX == length-1 && currY == length-1 {
			return grid[currX][currY]
		}
		value := grid[currX][currY]
		for _, move := range listOfMoves {
			x := currX + move[0]
			y := currY + move[1]
			if x < 0 || y < 0 || x >= length || y >= length || grid[x][y] != 0 {
				continue
			}
			grid[x][y] = value + 1
			queue = append(queue, [2]int{x, y})
		}
	}

	return -1
}
