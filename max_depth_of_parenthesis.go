package dsa

func maxDepth(s string) int {
	maxD := 0
	currD := 0
	for _, e := range s {
		if e == '(' {
			currD++
			maxD = max(currD, maxD)
		} else if e == ')' {
			currD--
		}
	}
	return maxD
}
