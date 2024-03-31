package dsa

import "slices"

// not done yet
func countSubarrays(nums []int, k int) int64 {
	var count int64 = 0
	maxE := slices.Max[[]int](nums)
	maxCount := 0
	currMaxCount := make([]int, len(nums))
	for i, e := range nums {
		if e == maxE {
			maxCount++
		}
		currMaxCount[i] = maxCount
	}

	j := 0
	for i, e := range currMaxCount {
		if e >= k {
			count += int64(i - j + 1)
			j = i
		}
	}
	return count
}
