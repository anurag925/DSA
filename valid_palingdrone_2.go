package dsa

func stringReverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}
func validPalindrome(s string) bool {
	count := 0
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] != s[end] {
			count++
			end--
		}
		start++
		end--
	}

	if count <= 1 {
		return true
	}

	count = 0
	start, end = 0, len(s)-1
	for start <= end {
		if s[start] != s[end] {
			count++
			start++
		}
		start++
		end--
	}
	if count == 1 {
		return true
	}

	return false
}
