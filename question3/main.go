package main

import (
	"fmt"
	"strings"
)

// radar, madam, naan
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Print("Input: ")
	fmt.Scan(&input)

	if isPalindrome(input) {
		fmt.Println("Output: True")
	} else {
		fmt.Println("Output: False")
	}
}
