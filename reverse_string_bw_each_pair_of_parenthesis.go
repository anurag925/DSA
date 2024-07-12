package dsa

func reverseParentheses(s string) string {
	temp := make([]string, 0)
	stack := make([]string, 0)
	for _, e := range s {
		if string(e) == ")" {
			for len(stack) != 0 && stack[len(stack)-1] != "(" {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(s)-1]
			stack = append(stack, temp...)
			temp = make([]string, 0)
		} else {
			stack = append(stack, string(e))
		}
	}
	res := ""
	for _, e := range stack {
		res += e
	}
	return res
}
