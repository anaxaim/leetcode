package main

import "fmt"

func isValid(s string) bool {
	parMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := make([]string, 0)

	if len(s) == 1 {
		return false
	}

	for _, v := range s {
		if string(v) == "(" || string(v) == "[" || string(v) == "{" {
			stack = append(stack, string(v))
		} else {
			if len(stack) == 0 {
				return false
			}

			if par, ok := parMap[string(v)]; ok && par != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("{[]}"))   // true
	fmt.Println(isValid("("))      // false
	fmt.Println(isValid("(("))     // false
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("]"))      // false
	fmt.Println(isValid("(()]"))   // false
}
