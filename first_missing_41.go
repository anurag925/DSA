package dsa

func firstMissingPositive(nums []int) int {
	l := len(nums)
	for _, v := range nums {
		ele := v
		for ele > 0 && ele <= l {
			p := nums[ele-1]
			nums[ele-1] = ele
			if p == ele {
				break
			}
			ele = p
		}
	}
	for i, v := range nums {
		if v <= 0 || (v != i+1) {
			return i + 1
		}
	}
	return l + 1
}
