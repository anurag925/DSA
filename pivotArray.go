package dsa

import "slices"

func pivotArray(nums []int, pivot int) []int {
	slices.SortStableFunc(nums, func(a, b int) int { return a - b })
	return nums
}
