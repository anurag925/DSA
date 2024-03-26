package dsa

import (
	"fmt"
	"slices"
)

func isOverlapp(a, b []int) {

}

func findMinArrowShots(points [][]int) int {
	count := 0
	if len(points) == 1 {
		return 1
	}
	slices.SortFunc[[][]int](points, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
	fmt.Println(points)
	i, j, n := 0, 1, len(points)
	minEnd := points[i][1]
	maxStart := points[i][0]
	for i < n && j < n {
		fmt.Println("check", i, j, maxStart, minEnd, points[i], points[j])
		if points[j][1] <= minEnd || points[j][0] >= maxStart {
			fmt.Println("overlap")
			if points[j][1] < minEnd {
				minEnd = points[j][1]
			}
			if points[j][0] > maxStart {
				maxStart = points[j][0]
			}
			j++
		} else {
			fmt.Println("no overlap")
			count++
			i = j
			minEnd = points[i][1]
			maxStart = points[i][0]
		}
	}
	return count
}
