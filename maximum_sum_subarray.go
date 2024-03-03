package dsa

func maxSubArray(nums []int) int {
	maxSum, currSum := 0, 0
	for _, i := range nums {
		currSum = max(currSum+i, i)
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}

func khadensAlgo(nums []int) int {
	maxSum, currSum := -9999999, 0
	for _, i := range nums {
		currSum += i
		if currSum > maxSum {
			maxSum = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}
