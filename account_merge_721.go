package dsa

import "fmt"

func accountsMerge(accounts [][]string) [][]string {
	res := make([][]string, 0)
	names := make(map[string]int)
	inames := make(map[int]string)
	emails := make(map[string]rune)
	remails := make(map[rune]string)
	emailCount := 'a'

	emailToName := make(map[rune][]int)
	for i, account := range accounts {
		for j, val := range account {
			if j == 0 {
				inames[i] = val
				names[val] = i
			} else {
				if _, ok := emails[val]; !ok {
					emails[val] = emailCount
					remails[emailCount] = val
					emailCount++
				}
				emailToName[emails[val]] = append(emailToName[emails[val]], i)
			}
		}
	}

	nameToName := make(map[int]int)
	for _, names := range emailToName {
		for _, name := range names {
			if _, ok := nameToName[name]; !ok {
				nameToName[name] = names[0]
			}
		}
	}
	fmt.Println(emailToName)
	fmt.Println(nameToName)
	emailName := make(map[rune]int)
	for email, names := range emailToName {
		for _, name := range names {
			emailName[email] = nameToName[name]
		}
	}

	fmt.Println(emailName)

	return res
}
