package dsa

import "slices"

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	for i := range heights {
		expected[i] = heights[i]
	}
	slices.Sort(expected)
	count := 0

	for i := range heights {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}
