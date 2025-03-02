package dsa

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
)

type FindClosestElementHeap struct {
	h []int
	x int
}

var _ heap.Interface = (*FindClosestElementHeap)(nil)

// Len implements heap.Interface.
func (h *FindClosestElementHeap) Len() int {
	return len(h.h)
}

// Less implements heap.Interface.
func (h *FindClosestElementHeap) Less(a int, b int) bool {
	amx := math.Abs(float64(a - h.x))
	bmx := math.Abs(float64(b - h.x))
	if amx < bmx {
		return true
	} else if amx == bmx {
		if a < b {
			return true
		}
	}
	return false
}

// Pop implements heap.Interface.
func (h *FindClosestElementHeap) Pop() any {
	top := h.h[h.Len()-1]
	h.h = h.h[:h.Len()-1]
	return top
}

// Push implements heap.Interface.
func (h *FindClosestElementHeap) Push(x any) {
	h.h = append(h.h, x.(int))
}

// Swap implements heap.Interface.
func (h *FindClosestElementHeap) Swap(i int, j int) {
	h.h[i], h.h[j] = h.h[j], h.h[i]
}

func findClosestElements(arr []int, k int, x int) []int {
	res := &FindClosestElementHeap{arr, x}
	heap.Init(res)
	for _, e := range arr {
		heap.Push(res, e)
		if res.Len() > k {
			heap.Pop(res)
		}
		fmt.Println(res.h)
	}
	return res.h
}
func findClosestElementsUsingSort(arr []int, k int, x int) []int {
	slices.SortFunc[[]int](arr, func(a, b int) int {
		amx := math.Abs(float64(a - x))
		bmx := math.Abs(float64(b - x))
		if amx < bmx {
			return -1
		} else if amx == bmx {
			if a < b {
				return -1
			}
		}
		return 1
	})
	res := arr[:k]
	slices.Sort(res)

	return res
}

type pointa struct {
	x, y int
	dis  float64
}

func kClosestUsingSort(points [][]int, k int) [][]int {
	res := make([]pointa, len(points))
	for i, v := range points {
		res[i] = pointa{
			x: v[0], y: v[1], dis: math.Sqrt(float64((v[0] * v[0]) + (v[1] * v[1]))),
		}
	}

	slices.SortFunc(res, func(a, b pointa) int { return int((a.dis - b.dis) * 100) })
	// fmt.Println(res)
	ans := make([][]int, 0)
	for _, v := range res[:k] {
		ans = append(ans, []int{v.x, v.y})
	}
	return ans
}
