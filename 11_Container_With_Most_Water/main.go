package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	mostWater := 0

	for left < right {
		square := min(height[left], height[right]) * (right - left)
		mostWater = max(square, mostWater)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return mostWater
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
}
