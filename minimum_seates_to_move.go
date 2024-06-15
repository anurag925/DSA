package dsa

import (
	"math"
	"slices"
)

func minMovesToSeat(seats []int, students []int) int {
	slices.Sort(seats)
	slices.Sort(students)
	count := 0

	for i := range students {
		count += int(math.Abs(float64(students[i] - seats[i])))
	}

	return count
}
