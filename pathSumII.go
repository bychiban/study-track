package main

import (
	"fmt"
)

func main() {
	fmt.Println(pathSum(&TreeNode{
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
				Val:   8,
				Right: &TreeNode{Val: 1},
			},
		}}, 122))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result, _ := pathSumCalc(root, targetSum)
	return result
}

func pathSumCalc(node *TreeNode, targetSum int) ([][]int, bool) {
	if node == nil {
		return nil, false
	}
	var result [][]int

	if targetSum-node.Val == 0 && node.Left == nil && node.Right == nil {
		result = append(result, []int{node.Val})
		return result, true
	}
	flag := false
	if leftResult, ok := pathSumCalc(node.Left, targetSum-node.Val); ok {
		for _, leftPath := range leftResult {
			result = append(result, append([]int{node.Val}, leftPath...))
		}
		flag = true
	}
	if rightResult, ok := pathSumCalc(node.Right, targetSum-node.Val); ok {
		for _, rightPath := range rightResult {
			result = append(result, append([]int{node.Val}, rightPath...))
		}
		flag = true
	}

	return result, flag
}
