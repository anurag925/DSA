package dsa

var dp = [37]int{0, 1, 1}
var done = false

func tribonacci(n int) int {
	if !done {
		for i := 3; i <= 37; i++ {
			dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		}
	}

	return dp[n]
}
