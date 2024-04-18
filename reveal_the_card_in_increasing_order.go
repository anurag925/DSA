package dsa

import "slices"

func deckRevealedIncreasing(deck []int) []int {
	slices.Sort(deck)
	q := make([]int, 0)
	for i := 0; i < len(deck); i++ {
		q = append(q, i)
	}

	ans := make([]int, len(deck))

	for _, card := range deck {
		top := q[0]
		q = q[1:]
		ans[top] = card
		if len(q) != 0 {
			top = q[0]
			q = append(q[1:], top)
		}
	}
	return ans
}
