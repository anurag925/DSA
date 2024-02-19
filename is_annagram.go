package dsa

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := make([]rune, 26)
	for _, e := range s {
		chars[e-97] += 1
	}
	for _, e := range t {
		chars[e-97] -= 1
	}
	for _, e := range chars {
		if e != 0 {
			return false
		}
	}
	return true
}
