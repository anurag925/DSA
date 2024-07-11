package dsa

import "fmt"

func findTheWinner(n int, k int) int {
	people := make([]int, n)
	for i := range people {
		people[i] = i + 1
	}

	z := n
	for z > 1 {
		fmt.Println(people)
		normalizer := z % k
		if normalizer == 0 {
			normalizer = z - 1
		}

		people = append(people[normalizer:], people[:normalizer+1]...)

		z = z - 1
	}

	return people[0]
}
