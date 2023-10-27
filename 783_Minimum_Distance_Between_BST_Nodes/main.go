package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	minDiff := math.MaxInt32
	prev := -1

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		if prev != -1 && root.Val-prev < minDiff {
			minDiff = root.Val - prev
		}
		prev = root.Val
		inorder(root.Right)
	}
	inorder(root)

	return minDiff
}

func main() {
	tr1 := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}
	fmt.Println(minDiffInBST(tr1)) // 1

	tr2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 49}}}
	fmt.Println(minDiffInBST(tr2)) // 1
}
