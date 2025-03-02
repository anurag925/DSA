package dsa

import "strings"

func backspaceCompare(s string, t string) bool {

	return strings.Compare(refactor(s), refactor(t)) == 0
}

func refactor(s string) string {
	i := 0
	for i >= 0 && i < len(s) {
		if s[i] == '#' {
			newi := max(i-2, 0)
			s = s[:newi] + s[i+1:]
		}
	}
	return s
}
