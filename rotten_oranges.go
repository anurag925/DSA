package dsa

import "fmt"

// func dfs(grid [][]byte, x, y int) {
// 	if x >= 0 && y >= 0 && x < rows && y < cols && grid[xd][yd] == '1' {

// 	}

// rows := len(grid)
// cols := len(grid[0])
// 	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
// for _, direction := range directions {
// 	xd, yd := (x + direction[0]), (y + direction[1])
// 	if xd >= 0 && yd >= 0 && xd < rows && yd < cols && grid[xd][yd] == '1' {
// 		grid[xd][yd] = '0'
// 		dfs(grid, xd, yd, )
// 	}
// }
// }

type point struct {
	x int
	y int
	m int
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// rottingGrid := make([][]int, len(grid))
	maxMin := -1
	queue := make([]point, 0)
	for i, col := range grid {
		for j, e := range col {
			if e == 0 {
				grid[i][j] = -1
			} else if e == 2 {
				grid[i][j] = 0
				queue = append(queue, point{i, j, 0})
			} else if e == 1 {
				grid[i][j] = 9999
			}
		}
	}

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			xd, yd, m := (top.x + direction[0]), (top.y + direction[1]), top.m+1
			if xd >= 0 && yd >= 0 && xd < rows && yd < cols && grid[xd][yd] != 0 && grid[xd][yd] != -1 && m < grid[xd][yd] {
				grid[xd][yd] = min(grid[xd][yd], m)
				queue = append(queue, point{xd, yd, m})
			}
		}
	}

	fmt.Println(grid)

	for _, col := range grid {
		for _, e := range col {
			if e > maxMin {
				maxMin = e
			}

			if e == 9999 {
				return -1
			}
		}
	}

	if maxMin == -1 {
		return 0
	}

	return maxMin
}
