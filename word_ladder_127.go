package dsa

func editDistance(a, b string) int {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	set := make(map[string]int)

	count := 0
	for _, e := range wordList {
		if _, ok := set[e]; !ok {
			set[e] = count
			count++
		}
	}
	if _, ok := set[endWord]; !ok {
		return 0
	}

	adjList := make([][]int, count)
	for _, v := range set {
		adjList[v] = make([]int, 0)
	}
	for i := range set {
		for j := range set {
			if i == j {
				continue
			}
			ed := editDistance(i, j)
			if ed == 1 {
				inti := set[i]
				intj := set[j]
				adjList[inti] = append(adjList[inti], intj)
			}
		}
	}

	dis := make([]int, count)
	for i, _ := range dis {
		dis[i] = 1e9
	}
	dis[set[beginWord]] = 0
	q := make([]int, 0)
	q = append(q, set[beginWord])

	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		for _, node := range adjList[top] {
			if dis[node] > (dis[top] + 1) {
				dis[node] = dis[top] + 1
				q = append(q, node)
			}
		}
	}

	for i, v := range dis {
		if i == set[endWord] && v >= 1e9 {
			return 0
		}
	}
	// +1 since we need to include the start word too
	return dis[set[endWord]] + 1
}

type wordDis struct {
	word string
	dis  int
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	set := make(map[string]bool)
	for _, e := range wordList {
		set[e] = true
	}
	if _, ok := set[endWord]; !ok {
		return 0
	}
	q := make([]wordDis, 0)
	q = append(q, wordDis{beginWord, 1})

	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		for i := 0; i < len(top.word); i++ {
			for j := 0; j < 26; j++ {
				newChar := string(byte(j + 'a'))
				newS := top.word[:i] + newChar + top.word[i+1:]
				if newS == endWord {
					return top.dis + 1
				}
				if _, ok := set[newS]; ok {
					q = append(q, wordDis{newS, top.dis + 1})
					delete(set, newS)
				}
			}
		}
	}
	return 0
}
