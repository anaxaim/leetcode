package main

import "fmt"

func rob(nums []int) int {
	lenNums := len(nums)

	dp := make([]int, lenNums+1)

	dp[1] = nums[0] // 0 1 0 0 0
	// dp[2] = max(1, 2 + 0) = 2
	// dp[3] = max(2, 3 + 1) = 4
	// dp[4] = max(4, 1 + 2) = 4
	for i := 2; i < len(nums)+1; i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}

	return dp[len(nums)]
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))    //4
	fmt.Println(rob([]int{2, 7, 9, 3, 1})) //12
}
