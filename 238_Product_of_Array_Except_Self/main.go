package main

import "fmt"

func productExceptSelf(nums []int) []int {
	productNums := make([]int, len(nums))

	// products to the left of each element -> [1, 2, 3, 4] -> [1, 1, 2, 6]
	// [1, 0, 0, 0] -> [1, 1*1, 0, 0] -> [1, 1, 2*1*1, 0] -> -> [1, 1, 2*1*1, 3*2*1*1]
	leftProduct := 1
	for i := 0; i < len(nums); i++ {
		productNums[i] = leftProduct
		leftProduct *= nums[i]
	}

	// products to the right of each element -> [1, 1, 2, 6] -> [24, 12, 8, 6]
	// [1, 1, 2, 6*1] -> [1, 1, 4*2, 6] -> [1, 12*1, 8, 6] -> -> [24*1, 12, 8, 6]
	rightProduct := 1
	for i := len(productNums) - 1; i >= 0; i-- {
		productNums[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return productNums
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      // [24, 12, 8, 6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // [0, 0, 9, 0, 0]
}
