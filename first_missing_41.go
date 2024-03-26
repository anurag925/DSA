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

func firstMissingPositive2(nums []int) int {
	checkMap := make(map[int]bool)
	for _, i := range nums {
		if i > 0 {
			checkMap[i] = true
		}
	}
	for i := 1; i < len(nums); i++ {
		if _, ok := checkMap[i]; !ok {
			return i
		}
	}
	return len(nums)
}
