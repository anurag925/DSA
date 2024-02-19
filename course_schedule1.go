package dsa

// detect a cycle in loop

func isCyclePresent(key int, adj map[int][]int, visited, recStack map[int]bool) bool {
	visited[key] = true
	recStack[key] = true

	for _, path := range adj[key] {
		if !visited[path] {
			if res := isCyclePresent(path, adj, visited, recStack); res {
				return true
			}
		} else if recStack[path] {
			return true
		}
	}
	recStack[key] = false
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	for _, e := range prerequisites {
		to, from := e[0], e[1]
		if _, ok := adj[from]; !ok {
			adj[from] = make([]int, 0)
		}
		adj[from] = append(adj[from], to)
	}

	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for key := range adj {
		if !visited[key] {
			if res := isCyclePresent(key, adj, visited, recStack); res {
				return false
			}
		}
	}

	return true
}
