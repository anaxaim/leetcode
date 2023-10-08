package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	if len(s) == len(t) {
		return strings.EqualFold(s, t)
	}

	counterS, counterT := 0, 0
	for counterS < len(s) && counterT < len(t) {
		if s[counterS] == t[counterT] {
			counterS++
		}
		counterT++
	}

	return counterS == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
	fmt.Println(isSubsequence("acb", "ahbgdc")) // false
	fmt.Println(isSubsequence("bb", "ahbgdc"))  // false
}
