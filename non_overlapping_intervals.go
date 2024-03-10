package dsa

import (
	"fmt"
	"slices"
)

// func isOverlap(a, b []int) bool {
// 	return a[0] <= b[1] && b[0] <= a[1]
// }

func eraseOverlapIntervals(intervals [][]int) int {
	count := 0
	slices.SortFunc[[][]int](intervals, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	n := len(intervals)

	fmt.Println("the array is ", intervals)
	// for i < n {
	// 	j := i + 1
	// 	if j < n && !isOverlap(intervals[i], intervals[j]) {
	// 		fmt.Println("overlap ", intervals[i], intervals[j], i, j)
	// 		count++
	// 	}
	// 	i = j
	// }
	for i := range intervals[1:] {
		j := i - 1
		if !isOverlap(intervals[i], intervals[j]) {
			count++
		}
	}

	return n - count
}
