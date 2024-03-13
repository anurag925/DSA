package dsa

import (
	"slices"
)

func customSortString(order string, s string) string {
	orderH := make(map[byte]int)
	for i := range order {
		orderH[order[i]] = i
	}
	sortingString := make([]byte, 0)
	nonSortingString := make([]byte, 0)
	for i := range s {
		if _, ok := orderH[s[i]]; ok {
			sortingString = append(sortingString, s[i])
		} else {
			nonSortingString = append(nonSortingString, s[i])
		}
	}

	slices.SortFunc[[]byte](sortingString, func(a, b byte) int {
		aval, aok := orderH[a]
		bval, bok := orderH[b]
		if aok && bok {
			return aval - bval
		}
		return int(a - b)
	})

	sortingString = append(sortingString, nonSortingString...)

	return string(sortingString[:])
}
