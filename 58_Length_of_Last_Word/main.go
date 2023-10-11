package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	strList := strings.Fields(s)

	return len(strList[len(strList)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))                 // 5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // 4
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       // 6
}
