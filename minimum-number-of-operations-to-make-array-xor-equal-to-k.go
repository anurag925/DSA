package dsa

import "strconv"

func minOperations(nums []int, k int) int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	count := 0
	bin1 := strconv.FormatInt(int64(xor), 2)
	bin2 := strconv.FormatInt(int64(k), 2)

	n := max(len(bin1), len(bin2))

	var a, b byte
	for i := 0; i < n; i++ {
		if i < len(bin1) {
			a = bin1[i]
		} else {
			a = byte(0)
		}

		if i < len(bin2) {
			b = bin2[i]
		} else {
			b = byte(0)
		}

		if a != b {
			count++
		}
	}
	return count
}
