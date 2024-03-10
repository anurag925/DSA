package dsa

import (
	"fmt"
	"strings"
)

// bad solve
func shortestSubstrings(arr []string) []string {
	ans := make([]string, len(arr))

	for i, e := range arr {
		res := ""
		for a := 0; a < len(e); a++ {
			for b := a + 1; b <= len(e); b++ {
				subStr := e[a:b]
				uniq := true
				for j, comp := range arr {
					fmt.Println(e, subStr, comp, i != j, uniq)
					if i != j && strings.Contains(comp, subStr) {
						uniq = false
						break
					}
				}
				if uniq {
					if res == "" || len(subStr) < len(res) || (len(subStr) == len(res) && subStr < res) {
						res = subStr
						continue
					}
				}
			}
		}
		ans[i] = res
	}

	return ans
}
