package dsa

func subarraySum(nums []int, k int) int {

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
