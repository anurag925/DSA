package dsa

func maxSubArray(nums []int) int {
	maxSum, currSum := 0, 0
	for _, i := range nums {
		currSum = max(currSum+i, i)
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}
