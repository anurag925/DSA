package dsa

import "fmt"

func isOverlapping(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func merge(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	merged := false
	i := 0
	for i < len(intervals) {
		for !merged && isOverlapping(newInterval, intervals[i]) {
			fmt.Println(i, intervals[i], newInterval)
			newInterval = merge(newInterval, intervals[i])
			merged = true
			i++
		}
		if merged && newInterval != nil {
			res = append(res, newInterval)
			newInterval = nil
		} else if newInterval[0] < intervals[i][0] {
			res = append(res, newInterval)
			merged = true
			newInterval = nil
		}
		fmt.Println(res)
		if i >= len(intervals) {
			break
		}
		res = append(res, intervals[i])
		i++
	}
	return res

}
