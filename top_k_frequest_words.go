package dsa

import (
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
