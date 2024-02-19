package dsa

import (
	"slices"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func kClosest(points [][]int, k int) [][]int {
	slices.SortFunc[[][]int](points, func(a, b []int) int {
		return (a[0]*a[0] + a[1]*a[1]) - (b[0]*b[0] + b[1]*b[1])
	})
	return points[:k]
}

func kClosestHeap(points [][]int, k int) [][]int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		i := a.([]int)
		j := b.([]int)
		return (i[0]*i[0] + i[1]*i[1]) - (j[0]*j[0] + j[1]*j[1])
	})
	for _, i := range points {
		pq.Enqueue(i)
	}
	res := make([][]int, k)
	for i := 0; i < k; i++ {
		j, _ := pq.Dequeue()
		res[i] = j.([]int)
	}
	return res
}
