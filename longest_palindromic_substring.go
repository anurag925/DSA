package dsa

func longestPalindrome(s string) string {
	res := ""
	for i := range s {
		// for odd length
		a, b := i, i
		for a >= 0 && b < len(s) && s[a] == s[b] {
			if len(s[a:b+1]) > len(res) {
				res = s[a : b+1]
			}
			a--
			b++
		}

		// for even length
		a, b = i, i+1
		for a >= 0 && b < len(s) && s[a] == s[b] {
			if len(s[a:b+1]) > len(res) {
				res = s[a : b+1]
			}
			a--
			b++
		}
	}
	return res
}
