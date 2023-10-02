package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	for _, v := range nums {
		if v != val {
			nums[count] = v
			count++
		}
	}
	return count
}

func main() {
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums1, 2)) // 5
	nums2 := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums2, 2)) // 2
}
