package dsa

import (
	"slices"
)

func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func mergeIntervals(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func mergeInterval(intervals [][]int) [][]int {
	slices.SortFunc[[][]int](intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	mergedInterval := make([][]int, 0)
	mergedInterval = append(mergedInterval, intervals[0])
	for i := 1; i < len(intervals); i++ {
		currInterval := mergedInterval[len(mergedInterval)-1]
		mergingInterval := intervals[i]
		if isOverlap(mergingInterval, currInterval) {
			newMergedInterval := mergeIntervals(mergingInterval, currInterval)
			mergedInterval[len(mergedInterval)-1] = newMergedInterval
		} else {
			mergedInterval = append(mergedInterval, mergingInterval)
		}
	}
	return mergedInterval
}
