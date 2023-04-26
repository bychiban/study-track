package main

import (
	"fmt"
)

func main() {
	fmt.Println(inorderTraversal(&TreeNode{
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
		}}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorderTraversalCalc(root, &result)
	return result
}

func inorderTraversalCalc(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorderTraversalCalc(node.Left, result)
	*result = append(*result, node.Val)
	inorderTraversalCalc(node.Right, result)
}
