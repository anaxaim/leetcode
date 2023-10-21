package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	}

	if x < 10 {
		return true
	}

	digits := 0
	// 121
	// digits = 1 | x = 12 -> 1 == x/10
	for x > 0 {
		digits = digits*10 + x%10
		x /= 10

		if digits == x || digits == x/10 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isPalindrome(121))   // true
	fmt.Println(isPalindrome(1221))  // true
	fmt.Println(isPalindrome(12521)) // true
	fmt.Println(isPalindrome(-121))  // false
	fmt.Println(isPalindrome(10))    // false
	fmt.Println(isPalindrome(100))   // false
}
