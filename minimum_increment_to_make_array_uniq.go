package dsa

import "slices"

func minIncrementForUnique(nums []int) int {
	count := 0
	maxNum := slices.Max(nums)
	arr := make([]int, maxNum+1)
	for _, e := range nums {
		arr[e]++
	}

	currCount := 0
	for i, e := range arr {
		if e == 1 {
			continue
		}
		if e == 0 && currCount > 0 {
			currCount--
			arr[i] = 1
			continue
		}

		currCount += e - 1
		arr[i] = 1
		count += e - 1
	}
	return count
}
