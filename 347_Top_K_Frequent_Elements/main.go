package main

import (
	"fmt"
	"sort"
)

type pair struct {
	val int
	cnt int
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}

	pairs := make([]pair, 0, len(counts))
	for key, val := range counts {
		pairs = append(pairs, pair{key, val})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].cnt > pairs[j].cnt
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i].val
	}

	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1, 2]
	fmt.Println(topKFrequent([]int{1}, 1))                // [1]
}
