package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasPathSum(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		}}, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if targetSum-root.Val == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	} else if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	return false
}
