package dsa

import (
	"fmt"
	"strings"
)

func removeOccurrences(s string, part string) string {
	res := ""
	for _, e := range s {
		res += string(e)
		fmt.Println(e, part[len(part)-1], res[len(res)-len(part):])
		if byte(e) == part[len(part)-1] && res[len(res)-len(part):] == part {
			fmt.Println(res[:len(res)-len(part)])
			res = res[:len(res)-len(part)]
		}
	}

	return res
}

func removeOccurrences2(s string, part string) string {
	for strings.Contains(s, part) {
		s = strings.Replace(s, part, "", 1)
	}
	return s
}
