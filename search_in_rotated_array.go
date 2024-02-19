package dsa

import "slices"

func searchLin(nums []int, target int) int {
	return slices.Index[[]int, int](nums, target)
}

func searchBinary(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		// left half of the array is sorted [4, 5, 6, 7, 8, 1, 2, 3]
		if nums[start] <= nums[mid] {
			if nums[start] <= target && nums[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { // right half of the array is sorted [6, 7, 8, 1, 2, 3, 4, 5]
			if nums[mid] < target && nums[end] >= target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
