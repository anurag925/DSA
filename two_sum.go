package dsa

import "github.com/emirpasic/gods/maps/hashmap"

func twoSum(nums []int, target int) []int {
	m := hashmap.New()
	for i, e := range nums {
		val, ok := m.Get(target - e)
		if ok {
			return []int{val.(int), i}
		}
		m.Put(e, i)
	}
	return []int{-1, -1}
}
