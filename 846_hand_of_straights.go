package dsa

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) == 0 || len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	cardCount := make(map[int]int)
	for _, card := range hand {
		cardCount[card]++
	}

	for _, card := range hand {
		if cardCount[card] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if cardCount[card+i] == 0 {
				return false
			}
			cardCount[card+i]--
		}
	}

	return true

}
