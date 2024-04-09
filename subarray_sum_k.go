package dsa

func subarraySum2(nums []int, k int) int {

	prefix := make([]int, len(nums))
	sumH := make(map[int]int)
	count := 0
	curr := 0
	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		prefix[i] = curr
		sumH[curr]++
		// curr := 0
		// for j := i; j < len(nums); j++ {
		// 	curr += nums[j]
		// 	if curr == k {
		// 		count++
		// 	}
		// }
	}
	return count
}

func subarraySum(nums []int, k int) int {
	count := 0
	prefix := make([]int, len(nums))
	hash := make(map[int]int)
	prefix[0] = nums[0]
	for i, e := range nums[1:] {
		prefix[i] = nums[i-1] + e
	}

	for _, e := range prefix {
		if e == k {
			count++
		}
		val, ok := hash[e-k]
		if ok {
			count += val
		}
		hash[e]++
	}

	return count
}
