package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, _ int) {
	for i := range nums1[m:] {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}

func main() {
	nums11 := []int{1, 2, 3, 0, 0, 0}
	merge(nums11, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums11)

	nums12 := []int{1}
	merge(nums12, 1, []int{}, 0)
	fmt.Println(nums12)

	nums13 := []int{0}
	merge(nums13, 0, []int{1}, 1)
	fmt.Println(nums13)
}
