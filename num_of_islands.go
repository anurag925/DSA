package dsa

func dfs(grid [][]byte, x, y int) {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, direction := range directions {
		xd, yd := (x + direction[0]), (y + direction[1])
		if xd >= 0 && yd >= 0 && xd < rows && yd < cols && grid[xd][yd] == '1' {
			grid[xd][yd] = '0'
			dfs(grid, xd, yd)
		}
	}
}

func numIslands(grid [][]byte) int {
	count := 0
	for i, col := range grid {
		for j, e := range col {
			if e == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}
