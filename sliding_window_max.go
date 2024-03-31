package dsa

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	j := 0
	s := make([]int, 0)
	for i, e := range nums {
		for len(s) > 0 && e > s[len(s)-1] {
			s = s[:len(s)-1]
		}
		s = append(s, e)
		if i-j+1 == k {
			top := s[0]
			res = append(res, top)
			if top == nums[j] {
				s = s[1:]
			}
			j++
		}
	}
	return res
}
