package dsa

func updateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	res := make([][]int, 0)
	stack := make([][]int, 0)

	for i, row := range mat {
		res = append(res, make([]int, len(row)))
		for j, v := range row {
			if v == 0 {
				res[i][j] = 0
				stack = append(stack, []int{i, j, 0})
			} else {
				res[i][j] = 999999
			}
		}
	}
	for len(stack) != 0 {
		curr := stack[0]
		x, y, dis := curr[0], curr[1], curr[2]
		stack = stack[1:]
		for _, d := range directions {
			xd, yd, disd := (x + d[0]), (y + d[1]), (dis + 1)
			if xd >= 0 && xd < rows && yd >= 0 && yd < cols && res[xd][yd] != 0 && res[xd][yd] > disd {
				stack = append(stack, []int{xd, yd, disd})
				res[xd][yd] = disd
			}
		}
	}
	return res
}
