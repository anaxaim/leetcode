package main

import "fmt"

func removeDuplicates(nums *[]int) int {
	count := 1
	for i := 1; i < len(*nums); i++ {
		if (*nums)[i-1] != (*nums)[i] {
			(*nums)[count] = (*nums)[i]
			count++
		}
	}

	*nums = (*nums)[:count]

	return count
}

func main() {
	nums1 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(&nums1)) // 2
	fmt.Println(nums1)                    // [1 2]

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} // 5
	fmt.Println(removeDuplicates(&nums2))
	fmt.Println(nums2) // [0 1 2 3 4]
}
