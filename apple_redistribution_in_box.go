package dsa

import (
	"slices"
)

func minimumBoxes(apple []int, capacity []int) int {
	allApples := 0
	for _, e := range apple {
		allApples += e
	}

	slices.SortFunc[[]int](capacity, func(a, b int) int {
		return b - a
	})

	for i, e := range capacity {
		allApples -= e
		if allApples <= 0 {
			return i + 1
		}
	}

	return -1
}
