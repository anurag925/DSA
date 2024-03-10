package dsa

import (
	"fmt"
	"slices"
)

func bagOfTokensScore(tokens []int, power int) int {
	maxScore := 0
	currScore := 0
	slices.Sort(tokens)
	start, end := 0, len(tokens)-1

	for start <= end {
		fmt.Println(tokens[start], tokens[end], power, currScore)
		if tokens[start] <= power {
			currScore++
			power -= tokens[start]
			start++
		} else if currScore > 1 {
			currScore--
			power += tokens[end]
			end--
		} else {
			fmt.Println("breaking")
			break
		}
		maxScore = max(maxScore, currScore)
	}

	return maxScore
}

// [100,200,300,400]
// [100, 200, 300, 400]
