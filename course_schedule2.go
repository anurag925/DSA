package dsa

// topological sort

type topoSort struct {
	visited     []bool
	adjList     [][]int
	parentStack []int
	recStack    []bool
}

// func (t *topoSort) topologicalDFSWithoutCycle(i int) {
// 	t.visited[i] = true
// 	for _, path := range t.adjList[i] {
// 		if !t.visited[path] {
// 			t.topologicalDFSWithoutCycle(path)
// 		}
// 	}
// 	t.parentStack = append(t.parentStack, i)
// }

func (t *topoSort) topologicalDFS(i int) bool {
	t.visited[i] = true
	t.recStack[i] = true
	for _, path := range t.adjList[i] {
		if !t.visited[path] && !t.topologicalDFS(path) {
			return false
		} else if t.recStack[path] {
			return false
		}
	}
	t.recStack[i] = false
	t.parentStack = append(t.parentStack, i)
	return true
}

func (t *topoSort) result() []int {
	return t.parentStack
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	recStack := make([]bool, numCourses)
	parentStack := make([]int, 0)
	for _, e := range prerequisites {
		from, to := e[0], e[1]
		adjList[from] = append(adjList[from], to)
	}
	topo := topoSort{visited, adjList, parentStack, recStack}

	for i := range adjList {
		if !topo.visited[i] {
			if !topo.topologicalDFS(i) {
				return []int{}
			}
		}
	}

	return topo.result()
}
