package main

import "fmt"

func main() {
	fmt.Println(isOverlap([]int{1, 4}, []int{2, 3}))
}
func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]

}
