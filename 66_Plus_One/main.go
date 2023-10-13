package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))    // [1, 2, 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // [4, 3, 2, 2]
	fmt.Println(plusOne([]int{9}))          // [1, 0]
	fmt.Println(plusOne([]int{9, 9}))       // [1, 0, 0]
}
