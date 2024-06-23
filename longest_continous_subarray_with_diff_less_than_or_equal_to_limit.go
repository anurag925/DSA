package dsa

import (
	"fmt"
	"math"
	"slices"
)

func longestSubarray(nums []int, limit int) int {
	count := 0
	minkeeper := make([]int, 0)
	i := 0
	for j := 0; j < len(nums); j++ {
		for len(minkeeper) > 0 && minkeeper[len(minkeeper)-1] > nums[j] {
			minkeeper = minkeeper[:len(minkeeper)-1]
		}
		minkeeper = append(minkeeper, nums[j])
		for math.Abs(float64(minkeeper[0]-nums[j])) > float64(limit) {
			fmt.Println(math.Abs(float64(minkeeper[0]-nums[j])), float64(limit))
			if minkeeper[0] == nums[i] {
				minkeeper = minkeeper[1:]
			}
			i++
		}
		if math.Abs(float64(minkeeper[0]-nums[j])) <= float64(limit) {
			count = max(count, j-i+1)
		}
		fmt.Println(minkeeper, i, j, count)
	}
	return count
}

func longestSubarray2(nums []int, limit int) int {
	count := 0
	i := 0
	for j := 0; j < len(nums); j++ {
		if slices.Max(nums[i:j+1])-slices.Min(nums[i:j+1]) <= limit {
			count = max(count, j-i+1)
		} else {
			i++
		}
	}
	return count
}

func longestSubarray3(nums []int, limit int) int {
	count := 0
	minkeeper := make([]int, 0)
	maxkeeper := -1
	maxi := -1
	i := 0
	for j := 0; j < len(nums); j++ {
		for len(minkeeper) > 0 && minkeeper[len(minkeeper)-1] > nums[j] {
			minkeeper = minkeeper[:len(minkeeper)-1]
		}
		minkeeper = append(minkeeper, nums[j])

		if maxkeeper < nums[j] || i > maxi {
			maxkeeper = nums[j]
			maxi = j
		}
		for math.Abs(float64(minkeeper[0]-maxkeeper)) > float64(limit) {
			fmt.Println(math.Abs(float64(minkeeper[0]-maxkeeper)), float64(limit))
			if minkeeper[0] == nums[i] {
				minkeeper = minkeeper[1:]
			}
			i++
			if maxkeeper < nums[j] || i > maxi {
				maxkeeper = nums[j]
				maxi = j
			}
		}
		if math.Abs(float64(minkeeper[0]-maxkeeper)) <= float64(limit) {
			count = max(count, j-i+1)
		}
		fmt.Println(minkeeper, i, j, count)
	}
	return count
}

func longestSubarray4(nums []int, limit int) int {
	count := 0
	minkeeper := make([]int, 0)
	maxkeeper := make([]int, 0)
	i := 0
	for j := 0; j < len(nums); j++ {
		for len(minkeeper) > 0 && minkeeper[len(minkeeper)-1] > nums[j] {
			minkeeper = minkeeper[:len(minkeeper)-1]
		}
		minkeeper = append(minkeeper, nums[j])

		for len(maxkeeper) > 0 && maxkeeper[len(maxkeeper)-1] < nums[j] {
			maxkeeper = maxkeeper[:len(maxkeeper)-1]
		}
		maxkeeper = append(maxkeeper, nums[j])

		for math.Abs(float64(minkeeper[0]-maxkeeper[0])) > float64(limit) {
			fmt.Println(math.Abs(float64(minkeeper[0]-nums[j])), float64(limit))
			if minkeeper[0] == nums[i] {
				minkeeper = minkeeper[1:]
			}

			if maxkeeper[0] == nums[i] {
				maxkeeper = maxkeeper[1:]
			}
			i++
		}
		// if math.Abs(float64(minkeeper[0]-nums[j])) <= float64(limit) {
		// 	count = max(count, j-i+1)
		// }
		count = max(count, j-i+1)
		fmt.Println(minkeeper, i, j, count)
	}
	return count
}
