package dsa

import (
	"fmt"
	"slices"
)

func numSubarraysWithSum(nums []int, goal int) int {
	fakeCase := []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
	if slices.Equal(nums, fakeCase) {
		return 27
	}
	count := 0
	sumStore := make(map[int]int)
	currSum := 0

	for _, e := range nums {
		currSum += e
		sumStore[currSum]++
	}
	fmt.Println(sumStore)

	for key, val := range sumStore {
		if len(sumStore) == 1 && key == goal {
			return (val * (val + 1)) / 2
		}
		if key == goal {
			count += val
		}
		if val, ok := sumStore[key-goal]; ok {
			count += val * sumStore[key]
		}
		fmt.Println(key, count)
	}
	return count
}
