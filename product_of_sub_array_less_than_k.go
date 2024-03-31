package dsa

// brute force
func numSubarrayProductLessThanK1(nums []int, k int) int {
	count := 0
	for i := range nums {
		p := 1
		for _, e := range nums[i:] {
			p *= e
			if p < k {
				count++
			} else {
				break
			}
		}
	}
	return count
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	p := 1
	i, j := 0, 0
	for i < len(nums) {
		if i < p {
			count++
		}
		for j < len(nums) && p < k {
			p *= nums[j]
			if p < k {
				count++
			} else {
				break
			}
			j++
		}
		if p >= k {
			p /= nums[i]
		}
		i++
	}
	return count
}
