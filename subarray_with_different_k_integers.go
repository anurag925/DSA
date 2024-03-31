package dsa

func subarraysWithKDistinct(nums []int, k int) int {
	count := 0
	for i := range nums {
		uniq := make(map[int]int)
		for _, je := range nums[i:] {
			uniq[je]++
			if len(uniq) == k {
				count++
			}
			if len(uniq) > k {
				break
			}
		}
	}
	return count
}

// func subarraysWithKDistinct2(nums []int, k int) int {
// 	count := 0
// 	i, j := 0, 0
// 	uniq := make(map[int]int)
// 	for i < len(nums) {

// 	}
// 	return count
// }

func subarraysWithKDistinct3(nums []int, k int) int {
	count := 0
	uniq := make(map[int]int)
	for i := range nums {
		for _, je := range nums[i:] {
			uniq[je]++
			if len(uniq) == k {
				count++
			}
			if len(uniq) > k {
				break
			}
		}
		clear(uniq)
	}
	return count
}
