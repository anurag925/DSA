package main

import (
	"dsa"
)

func main() {
	// fmt.Println(isOverlap([]int{1, 4}, []int{2, 3}))
	dsa.MakeGood("ss")

	// res := "3" >= "30"
	// fmt.Print(res)
}
func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]

}
