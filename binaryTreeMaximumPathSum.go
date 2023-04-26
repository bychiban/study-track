package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxPathSum(&TreeNode{
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

func maxPathSum(root *TreeNode) int {
	result := -1 << 63
	maxPathSumCalc(root, &result)
	return result
}

func maxPathSumCalc(node *TreeNode, result *int) int {
	if node == nil {
		return 0
	}

	leftMaxPath := max(maxPathSumCalc(node.Left, result), 0)
	rigthMaxPath := max(maxPathSumCalc(node.Right, result), 0)
	*result = max(*result, leftMaxPath+rigthMaxPath+node.Val)
	return max(leftMaxPath, rigthMaxPath) + node.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
