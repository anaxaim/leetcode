package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{1, 2}
	}

	hash := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if index, ok := hash[diff]; ok {
			return []int{index + 1, i + 1}
		}
		hash[num] = i
	}

	return nil
}

func main() {
	res1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res1) // [1, 2]

	res2 := twoSum([]int{2, 3, 4}, 6)
	fmt.Println(res2) // [1, 3]

	res3 := twoSum([]int{3, 3}, 6)
	fmt.Println(res3) // [1, 2]

	res4 := twoSum([]int{2, 3, 3}, 6)
	fmt.Println(res4) // [2, 3]

	res5 := twoSum([]int{-1, 0}, -1)
	fmt.Println(res5) // [1, 2]
}
