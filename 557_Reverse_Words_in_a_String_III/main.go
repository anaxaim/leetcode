package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	var builder strings.Builder

	for i, w := range words {
		runes := []rune(w)
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		builder.WriteString(string(runes))
		if i < len(words)-1 {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest")) // s'teL ekat edoCteeL tsetnoc
	fmt.Println(reverseWords("God Ding"))                    // doG gniD
}
