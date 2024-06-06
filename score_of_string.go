package dsa

import "math"

func scoreOfString(s string) int {
	sum := 0
	var prev rune = 0
	for _, e := range s {
		if prev == 0 {
			prev = e
			continue
		}
		sum += int(math.Abs(float64(e - prev)))
		prev = e
	}
	return sum
}
