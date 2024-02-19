package dsa

import (
	"fmt"
)

type wordDisPath struct {
	word string
	path []string
	vis  []bool
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	set := make(map[string]int)
	vis := make([]bool, len(wordList))
	for i, e := range wordList {
		set[e] = i
	}
	if _, ok := set[endWord]; !ok {
		return [][]string{}
	}

	delete(set, endWord)
	q := make([]wordDisPath, 0)
	vis[set[beginWord]] = true
	pathBegin := wordDisPath{beginWord, []string{beginWord}, vis}
	q = append(q, pathBegin)

	res := make([][]string, 0)
	var maxLen int = 1e9

	for len(q) != 0 {
		// fmt.Println(q)
		top := q[0]
		q = q[1:]

		for i := 0; i < len(top.word); i++ {
			for j := 0; j < 26; j++ {
				newChar := string(byte(j + 'a'))
				newS := top.word[:i] + newChar + top.word[i+1:]
				fmt.Print(newS, ",")
				if newS == endWord {
					top.path = append(top.path, newS)
					if len(top.path) > maxLen {
						continue
					}
					if len(top.path) < maxLen {
						res = make([][]string, 0)
						maxLen = len(top.path)
					}
					res = append(res, top.path)
				} else if _, ok := set[newS]; ok && !top.vis[set[newS]] {
					pathCopy := make([]string, len(top.path))
					copy(pathCopy, top.path)
					pathCopy = append(pathCopy, newS)

					visCopy := make([]bool, len(top.vis))
					copy(visCopy, top.vis)
					visCopy[set[newS]] = true

					q = append(q, wordDisPath{newS, pathCopy, visCopy})
					// delete(set, newS)
				}
			}
			fmt.Println()
		}
	}
	return res
}
