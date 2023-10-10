package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for _, v := range magazine {
		m[v]++
	}

	for _, v := range ransomNote {
		if count, ok := m[v]; ok && count > 0 {
			m[v]--
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "ab"))  // false
	fmt.Println(canConstruct("aa", "aab")) // true
}
