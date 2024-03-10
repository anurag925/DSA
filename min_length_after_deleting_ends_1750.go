package dsa

func minimumLength(s string) int {
	n := len(s)
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			break
		}
		for (start+1) < n && s[start] == s[start+1] {
			start++
		}
		for (end-1) >= 0 && s[end] == s[end-1] {
			end--
		}
		start++
		end--
	}
	res := end - start + 1
	if res < 0 {
		return 0
	}
	return res
}
