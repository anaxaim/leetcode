package main

import "fmt"

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var median int

	for low <= high {
		median = low + (high-low)/2

		if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if nums[median] == target {
		return median
	}

	return low
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0)) // 0
	fmt.Println(searchInsert([]int{1, 2}, 2))       // 1
}
