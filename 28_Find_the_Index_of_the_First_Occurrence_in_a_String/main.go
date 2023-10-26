package main

import (
	"fmt"
	"strings"
)

// mysolution - O(nХm)
func strStr2(haystack string, needle string) int {
	splitted := strings.Split(haystack, needle)
	if len(splitted) > 1 {
		return len(splitted[0])
	}
	return -1
}

// Longest Prefix Suffix
func computeLPS(needle string, needleLen int, lps []int) {
	length := 0
	i := 1
	for i < needleLen {
		if needle[i] == needle[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
}

// KMP алгоритм Кнута-Морриса-Пратта
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen == 0 {
		return 0
	}

	lps := make([]int, needleLen)
	computeLPS(needle, needleLen, lps)

	for haystackIdx, needleIdx := 0, 0; haystackIdx < haystackLen; {
		if needle[needleIdx] == haystack[haystackIdx] {
			haystackIdx++
			needleIdx++
		}

		if needleIdx == needleLen {
			return haystackIdx - needleIdx
		} else if haystackIdx < haystackLen && needle[needleIdx] != haystack[haystackIdx] {
			if needleIdx != 0 {
				needleIdx = lps[needleIdx-1]
			} else {
				haystackIdx = haystackIdx + 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))     // 0
	fmt.Println(strStr("leetcode", "leeto"))    // -1
	fmt.Println(strStr("mississippi", "issip")) // 4
}
