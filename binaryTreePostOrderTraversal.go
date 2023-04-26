package main

import (
	"fmt"
)

func main() {
	fmt.Println(postorderTraversal(&TreeNode{
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

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorderTraversalCalc(root, &result)
	return result
}

func postorderTraversalCalc(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	postorderTraversalCalc(node.Left, result)
	postorderTraversalCalc(node.Right, result)
	*result = append(*result, node.Val)
}
