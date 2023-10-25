package main

// not in-place implementations
func rotate2(nums []int, k int) {
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		nums = append([]int{last}, nums[:len(nums)-1]...)
	}
}

func rotate(nums []int, k int) {
	// for cases when k >= len(nums)
	k = k % len(nums)

	for start, end := 0, len(nums)-1; start <= end; {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	for start, end := 0, k-1; start <= end; {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

	for start, end := k, len(nums)-1; start <= end; {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3) // [5 ,6, 7, 1, 2, 3, 4]
	rotate([]int{-1, -100, 3, 99}, 2)     // [3, 99, -1, -100]
	rotate([]int{-1}, 2)                  // [-1]
}
