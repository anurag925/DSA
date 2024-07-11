package dsa

import "fmt"

func minOperationsLog(logs []string) int {
	s := make([]string, 0)
	for _, e := range logs {
		fmt.Printf("the stack is %+v", s)
		if e == "../" {
			if len(s) != 0 {
				s = s[:len(s)-1]
			}

		} else if e == "./" {
			// s = s[:len(s)-1]
		} else {
			s = append(s, e)
		}
	}
	return len(s)
}
