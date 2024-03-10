package dsa

func divide(dividend int, divisor int) int {
	res := dividend / divisor
	if res > (2147483647) {
		return 2147483647
	} else if res < (-2147483648) {
		return -2147483648
	}

	return res
}
