package main

import "fmt"

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
}

func shortestBridge(grid [][]int) int {
	cellQueue := [][]int{}
	n := len(grid)
	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				isleSearch(&grid, &cellQueue, i, j)
				found = true
				break
			}
		}
	}
	distance := 0
	for {
		newQueue := [][]int{}
		for _, cell := range cellQueue {
			x, y := cell[0], cell[1]
			moves := [4][2]int{{x, y - 1}, {x, y + 1}, {x - 1, y}, {x + 1, y}}
			for _, move := range moves {
				x, y = move[0], move[1]
				if x >= 0 && y >= 0 && x < n && y < n {
					if (grid)[x][y] == 1 {
						return distance
					} else if (grid)[x][y] == 0 {
						grid[x][y] = -1
						newQueue = append(newQueue, []int{x, y})
					}
				}
			}
		}
		cellQueue = newQueue
		distance++
	}
}

func isleSearch(grid *[][]int, isle *[][]int, x, y int) {
	n := len(*grid)
	(*grid)[x][y] = 2
	(*isle) = append((*isle), []int{x, y})
	moves := [4][2]int{{x, y - 1}, {x, y + 1}, {x - 1, y}, {x + 1, y}}

	for _, move := range moves {
		x, y = move[0], move[1]
		if x >= 0 && y >= 0 && x < n && y < n && (*grid)[x][y] == 1 {
			isleSearch(grid, isle, x, y)
		}
	}
}
