package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")

	length := len(s)
	for i := 0; i < length/2; i++ {
		if !strings.EqualFold(string(s[i]), string(s[length-1-i])) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}
