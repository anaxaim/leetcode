package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, v := range t {
		m[v]++
	}

	for _, v := range s {
		if count, ok := m[v]; ok && count > 0 {
			m[v]--
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println(isAnagram("a", "ab"))            // false
}
