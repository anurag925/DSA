package dsa

func timeRequiredToBuy2(tickets []int, k int) int {

	count := 0
	n := len(tickets)
	i := 0
	for tickets[k] != 0 {
		for tickets[i] == 0 {
			i++
			if i >= n {
				i = 0
			}
		}
		tickets[i]--
		count++
		i++
		if i >= n {
			i = 0
		}
	}

	return count
}

func timeRequiredToBuy(tickets []int, k int) int {

	count := 0
	for i, e := range tickets {
		if i <= k {
			count += e
		} else {
			count += tickets[k]
		}
	}
	return count
}
