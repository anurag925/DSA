package dsa

// longest subarray with equal number of 0 and 1
func findMaxLength(nums []int) int {
	length := 0
	isum := make(map[int]int)
	isum[0] = -1
	sum := 0
	for i, e := range nums {
		if e == 0 {
			sum += -1
		} else {
			sum += 1
		}
		if val, ok := isum[sum]; ok {
			if (i - val) > length {
				length = i - val
			}
		} else {
			isum[sum] = i
		}
	}
	return length
}
