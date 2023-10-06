package main

import "fmt"

func removeDuplicates(nums *[]int) int {
	lenNums := len(*nums)

	switch lenNums {
	case 1, 2:
		return lenNums
	default:
		count := 2
		for i := 2; i < lenNums; i++ {
			if (*nums)[count-2] != (*nums)[i] {
				(*nums)[count] = (*nums)[i]
				count++
			}
		}
		*nums = (*nums)[:count]

		return count
	}
}

func main() {
	nums1 := []int{1}
	fmt.Println(removeDuplicates(&nums1)) // 1
	fmt.Println(nums1)                    // [1]

	nums2 := []int{0, 1}
	fmt.Println(removeDuplicates(&nums2)) // 2
	fmt.Println(nums2)                    // [0, 1]

	nums4 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(&nums4)) // 5
	fmt.Println(nums4)                    // [1, 1, 2, 2, 3]

	nums5 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(&nums5)) // 7
	fmt.Println(nums5)                    // [0, 0, 1, 1, 2, 3, 3]
}
