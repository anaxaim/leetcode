package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	windowSum := 0
	windowSize := math.MaxInt32

	start := 0
	end := 0
	for end < len(nums) {
		windowSum = windowSum + nums[end]
		for windowSum >= target {
			minWindowSize := (end + 1) - start
			if windowSize > minWindowSize {
				windowSize = minWindowSize
			}
			windowSum = windowSum - nums[start]
			start++
		}
		end++
	}
	if windowSize == math.MaxInt32 {
		return 0
	}

	return windowSize
}

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))          // 3
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
}
