package dsa

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		lastChar := stack[len(stack)-1]
		if (lastChar == '(' && char == ')') || (lastChar == '{' && char == '}') || (lastChar == '[' && char == ']') {
			stack = stack[:len(stack)-1]
			continue
		}
		break
	}
	return len(stack) == 0
}
