package dsa

import (
	"slices"
	"strconv"
)

func largestNumber(nums []int) string {
	numss := make([]string, len(nums))
	for i, e := range nums {
		numss[i] = strconv.Itoa(e)
	}
	slices.SortFunc(numss, func(a, b string) int {
		one, _ := strconv.Atoi(a + b)
		two, _ := strconv.Atoi(b + a)
		return two - one
	})
	res := ""
	for _, e := range numss {
		if res == "0" || res == "" && e == "0" {
			res = "0"
			break
		}
		res += e
	}
	return res
}
