package dsa

import "slices"

func maximumHappinessSum(happiness []int, k int) int64 {
	slices.SortFunc[[]int](happiness, func(a, b int) int { return b - a })

	var mhappy int64 = 0
	for i := 0; i < k; i++ {
		if happiness[i]-i > 0 {
			mhappy += int64(happiness[i] - i)
		}

	}

	return mhappy
}
