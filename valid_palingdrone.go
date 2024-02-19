package dsa

import (
	"fmt"
	"strings"
	"unicode"
)

func PrintRuneOfAlphaNumeric() {
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i, c := range a {
		fmt.Printf("for char no %d val is %v ascii is %d \n", i, c, c)
	}
	// data := rune(a[0])
	fmt.Print(rune('0'), rune('P'))

}

func isPalindrome(s string) bool {
	n := len(s)
	start, end := 0, n-1
	for start <= end {
		startChar := rune(s[start])
		endChar := rune(s[end])
		if isAlpha(startChar) && isAlpha(endChar) {
			diff := startChar - endChar
			if diff != 32 && diff != 0 && diff != -32 {
				// fmt.Print(start, end, startChar, endChar, diff)
				return false
			}
			start += 1
			end -= 1
		} else if isNumeric(startChar) && isNumeric(endChar) {
			diff := startChar - endChar
			if diff != 0 {
				// fmt.Print(start, end, startChar, endChar, diff)
				return false
			}
			start += 1
			end -= 1
		} else if isAlphanumeric(startChar) && isAlphanumeric(endChar) {
			return false
		} else if isAlphanumeric(startChar) {
			end -= 1
		} else {
			start += 1
		}
	}
	return true
}

func isAlpha(i rune) bool {
	return (i >= 65 && i <= 90) || (i >= 97 && i <= 122)
}

func isNumeric(i rune) bool {
	return (i >= 48 && i <= 57)
}

func isAlphanumeric(i rune) bool {
	return isNumeric(i) || isAlpha(i)
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	clean_s := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			clean_s = append(clean_s, s[i])
		}
	}

	fmt.Println(clean_s)
	for i := 0; i < len(clean_s); i++ {
		if clean_s[i] != clean_s[len(clean_s)-1-i] {
			return false
		}
	}
	return true
}
