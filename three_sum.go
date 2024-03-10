package dsa

import (
	"slices"

	"github.com/emirpasic/gods/maps/hashmap"
)

func threeSum(nums []int) [][]int {
	// h := hashset.New()
	h := make([][]int, 0)
	res := hashmap.New()

	for i, e := range nums {
		res.Put(e, i)
	}

	for i, e := range nums {
		for j, e2 := range nums {
			if i != j {
				val, ok := res.Get(-(e + e2))
				if ok && val != i && val != j {
					arr := []int{e, e2, -(e + e2)}
					slices.Sort(arr)
					// h.Add(arr)
					h = append(h, arr)
				}
			}
		}
	}

	// resArr := make([][]int, 0)

	// for _, val := range h.Values() {
	// 	resArr = append(resArr, val.([]int))
	// }

	// return resArr
	return h
}
