package dsa

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	eod := 23*60 + 59 + 1
	minutes := make([]int, len(timePoints))
	res := math.MaxInt
	for i, e := range timePoints {
		timeSplit := strings.Split(e, ":")
		hour, _ := strconv.Atoi(timeSplit[0])
		minute, _ := strconv.Atoi(timeSplit[1])
		minutes[i] = hour*60 + minute
	}
	slices.Sort(minutes)
	for i, e := range minutes {
		if i+1 < len(minutes) {
			diff := minutes[i+1] - e
			if diff < res {
				res = diff
			}
		}
	}

	circularDiff := eod - minutes[len(minutes)-1] + minutes[0]
	if circularDiff < res {
		res = circularDiff
	}

	return res
}
