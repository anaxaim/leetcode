package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	// split string by words
	sSlice := strings.Fields(s)

	if len(sSlice) != len(pattern) {
		return false
	}

	// create two maps runeToWordMap = {'a': dog, 'b': cat} and otherwise
	runeToWordMap := make(map[rune]string)
	wordToRuneMap := make(map[string]rune)

	for i, v := range pattern {
		word := sSlice[i]
		// add letter from pattern as a key and word from s as a value
		if mV, ok := runeToWordMap[v]; ok {
			// if word in the same index is not equal to letter mapping -> it's wrong pattern
			if mV != word {
				return false
			}
		} else {
			runeToWordMap[v] = word
		}

		// add word as a key to the second map and letter as a value
		// if word has already in map and word to equal to current letter from pattern -> it's wrong pattern
		if ch, ok := wordToRuneMap[word]; ok && ch != v {
			return false
		}
		wordToRuneMap[word] = v
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))  // true
	fmt.Println(wordPattern("abba", "dog dog dog dog"))  // false
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  // false
}
