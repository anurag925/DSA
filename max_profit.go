package dsa

import "fmt"

func maxProfit(prices []int) int {
	suffix := make([]int, len(prices))
	prefix := make([]int, len(prices))

	last := 999999
	for i, e := range prices {
		if e < last {
			last = e
		}
		suffix[i] = e
	}

	last = -999999
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > last {
			last = prices[i]
		}
		prefix[i] = last
	}

	fmt.Print(suffix)
	fmt.Print(prefix)

	profit := 0
	for i := 0; i < len(prices); i++ {
		newProfit := prefix[i] - suffix[i]
		if newProfit > profit {
			profit = newProfit
		}
	}

	return profit
}
