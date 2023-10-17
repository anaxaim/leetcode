package main

import "fmt"

func maxProfit(prices []int) int {
	maxProf, curProf := 0, 0

	for i := 1; i < len(prices); i++ {
		curProf += prices[i] - prices[i-1]
		if curProf < 0 {
			curProf = 0
		}

		maxProf = max(maxProf, curProf)
	}

	return maxProf
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}
