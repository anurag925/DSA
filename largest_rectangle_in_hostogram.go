package dsa

import "math"

func largestRectangleArea(heights []int) int {
	area := 0
	smallerOnRight := nextSmallerOnRight(heights)
	smallerOnLeft := nextSmallerOnLeft(heights)
	for i, height := range heights {
		breadth := math.Abs(float64(smallerOnLeft[i])-float64(smallerOnRight[i])) - 1
		area = max(area, height*int(breadth))
	}
	return area
}

func nextSmallerOnRight(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	s := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && heights[s[len(s)-1]] >= heights[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			ans[i] = -1
		} else {
			ans[i] = s[len(s)-1]
		}
		s = append(s, i)
	}
	return ans
}

func nextSmallerOnLeft(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(s) > 0 && heights[s[len(s)-1]] >= heights[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			ans[i] = -1
		} else {
			ans[i] = s[len(s)-1]
		}
		s = append(s, i)
	}
	return ans
}
