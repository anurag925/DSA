package dsa

func Reverse(input string) (output string) {
	for _, r := range input {
		output = string(r) + output
	}

	return
}
func isPalindrome5(head *ListNode) bool {
	num := ""
	for head != nil {
		num = num + string(head.Val)
		head = head.Next
	}

	return num == Reverse(num)
}
