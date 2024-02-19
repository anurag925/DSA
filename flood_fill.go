package dsa

func isDirectionSafe(row, col, x, y int) bool {
	return x >= 0 && x < row && y >= 0 && y < row
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	row, col := len(image), len(image[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	stack := make([][]int, 0)
	stack = append(stack, []int{sr, sc})
	for len(stack) > 0 {
		x, y := stack[0][0], stack[0][1]
		stack = stack[1:]
		if image[x][y] == color {
			continue
		}
		image[x][y] = color
		for _, d := range directions {
			xd, yd := x+d[0], y+d[0]
			if isDirectionSafe(row, col, xd, yd) && image[xd][yd] != color {
				stack = append(stack, []int{xd, yd})
			}
		}
	}

	return image
}
