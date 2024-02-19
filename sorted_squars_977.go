package dsa

import "slices"

func sortedSquares(nums []int) []int {
	for i, e := range nums {
		nums[i] = e * e
	}
	slices.Sort[[]int](nums)
	return nums
}
