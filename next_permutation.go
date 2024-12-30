package dsa

import "slices"

func nextPermutation(nums []int) {
	n := len(nums)
	i, j := n-1, n-2
	for j >= 0 && i > 0 && nums[j] >= nums[i] {
		j--
		i--
	}
	if j == -1 {
		slices.SortStableFunc(nums, func(a, b int) int { return a - b })
		return
	}
	closest := 999
	closestindex := -1
	for i := j; i < n; i++ {
		diff := nums[i] - nums[j]
		if diff > 0 && diff < closest {
			closest = diff
			closestindex = i
		}
	}
	nums[j], nums[closestindex] = nums[closestindex], nums[j]
	// fmt.Println(nums, closestindex, j)
	slices.SortStableFunc(nums[j+1:], func(a, b int) int { return a - b })
}
