package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	freq := 1
	majority := nums[0]
	for k, v := range m {
		if v > freq {
			freq = v
			majority = k
		}
	}

	return majority
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))             //3
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) //2
}
