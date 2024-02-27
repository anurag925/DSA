package dsa

import (
	"container/heap"
	"slices"
	"strings"
)

type Wordncount struct {
	Word  string
	Count int
}

func topKFrequent(words []string, k int) []string {
	counter := make(map[string]int)
	for _, word := range words {
		counter[word] += 1
	}
	arr := make([]Wordncount, 0)
	for key, val := range counter {
		arr = append(arr, Wordncount{key, val})
	}
	slices.SortFunc[[]Wordncount, Wordncount](arr, func(a, b Wordncount) int {
		diff := b.Count - a.Count
		if diff == 0 {
			return strings.Compare(a.Word, b.Word)
		}
		return diff
	})

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i].Word
	}
	return res
}

type WordCount struct {
	w string
	c int
}

type WordCountHeap struct {
	h []WordCount
}

var _ heap.Interface = (*WordCountHeap)(nil)

// Len implements heap.Interface.
func (t *WordCountHeap) Len() int {
	return len(t.h)
}

// Less implements heap.Interface.
func (t *WordCountHeap) Less(i int, j int) bool {
	return t.h[i].c < t.h[j].c
}

// Pop implements heap.Interface.
func (t *WordCountHeap) Pop() any {
	res := t.h[t.Len()-1]
	t.h = t.h[:t.Len()-1]
	return res
}

// Push implements heap.Interface.
func (t *WordCountHeap) Push(x any) {
	t.h = append(t.h, x.(WordCount))
}

// Swap implements heap.Interface.
func (t *WordCountHeap) Swap(i int, j int) {
	t.h[i], t.h[j] = t.h[j], t.h[i]
}

func topKFrequent2(words []string, k int) []string {
	wordCountHash := make(map[string]int)
	wordCountHeap := &WordCountHeap{
		h: make([]WordCount, 0),
	}

	for _, i := range words {
		if val, ok := wordCountHash[i]; !ok {
			wordCountHash[i] = 1
		} else {
			wordCountHash[i] = val + 1
		}
	}

	heap.Init(wordCountHeap)

	for word, count := range wordCountHash {
		heap.Push(wordCountHeap, WordCount{word, count})
		if wordCountHeap.Len() == k {
			heap.Pop(wordCountHeap)
		}
	}

	res := make([]string, 0)

	for _, he := range wordCountHeap.h {
		res = append(res, he.w)
	}
	return res
}
