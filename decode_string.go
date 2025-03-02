package dsa

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	st := make([]string, 0)
	i := 0
	for i < len(s) {
		temp := ""
		for s[i] >= '0' && s[i] <= '9' {
			temp += string(s[i])
			i++
		}
		if temp != "" {
			st = append(st, temp)
			temp = ""
			continue
		}
		for s[i] >= 'a' && s[i] <= 'z' {
			temp += string(s[i])
			i++
		}
		if temp != "" {
			st = append(st, temp)
			temp = ""
			continue
		}
		if s[i] == '[' {
			i++
			continue
		}

		// when i == ']'
		text := st[len(st)-1]
		st = st[:len(st)-1]
		num := st[len(st)-1]
		st = st[:len(st)-1]
		numInt, err := strconv.Atoi(num)
		if err != nil {
			temp = num + text
		} else {
			for a := 0; a < numInt; a++ {
				temp += text
			}

		}
		st = append(st, temp)
		temp = ""
		i++
	}

	return strings.Join(st, "")
}
