package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	// 3,9,20,null,null,15,7
	tr1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(maxDepth(tr1)) // 3

	// [1,null,2]
	tr2 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(maxDepth(tr2)) // 2

	tr3 := &TreeNode{}
	fmt.Println(maxDepth(tr3)) // 1

}
