package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	count := 1
	for ; ; count++ {
		if count*count > x {
			return count - 1
		}
	}
}

func main() {
	fmt.Println(mySqrt(0)) // 0
	fmt.Println(mySqrt(2)) // 1
	fmt.Println(mySqrt(3)) // 1
	fmt.Println(mySqrt(4)) // 2
	fmt.Println(mySqrt(8)) // 2
}
