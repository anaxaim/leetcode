package main

import "fmt"

func climbStairsRecursive(n int) int {
	if n < 3 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	current := 0
	before := 3
	beforeBefore := 2
	for i := 4; i <= n; i++ {
		current = before + beforeBefore
		beforeBefore = before
		before = current
	}

	return current
}

func main() {
	fmt.Println(climbStairsRecursive(2))  // 2
	fmt.Println(climbStairsRecursive(3))  // 3
	fmt.Println(climbStairsRecursive(5))  // 8
	fmt.Println(climbStairsRecursive(6))  // 13
	fmt.Println(climbStairsRecursive(45)) // 1836311903
	fmt.Println(climbStairs(2))           // 2
	fmt.Println(climbStairs(3))           // 3
	fmt.Println(climbStairs(5))           // 8
	fmt.Println(climbStairs(6))           // 13
	fmt.Println(climbStairs(45))          // 1836311903
}
