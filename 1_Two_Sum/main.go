package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	hash := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if index, ok := hash[diff]; ok {
			return []int{index, i}
		}
		hash[num] = i
	}

	return nil
}

func main() {
	res1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res1) // [0, 1]

	res2 := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(res2) // [1, 2]

	res3 := twoSum([]int{3, 3}, 6)
	fmt.Println(res3) // [0, 1]

	res4 := twoSum([]int{3, 2, 3}, 6)
	fmt.Println(res4) // [0, 2]
}
