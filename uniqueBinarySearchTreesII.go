package main

import (
	"fmt"
)

func main() {
	printTrees(generateTrees(2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}
	return generateTreesCalc(list)
}

func generateTreesCalc(list []int) []*TreeNode {
	length := len(list)
	if length == 0 {
		return []*TreeNode{nil}
	}
	if length == 1 {
		return []*TreeNode{{Val: list[0]}}
	}
	result := []*TreeNode{}
	for i := 0; i < length; i++ {
		leftBranches := generateTreesCalc(list[:i])
		rightBranches := generateTreesCalc(list[i+1 : length])
		for _, leftBranch := range leftBranches {
			for _, rightBranch := range rightBranches {
				current := &TreeNode{Val: list[i], Left: leftBranch, Right: rightBranch}
				result = append(result, current)
			}
		}
	}
	return result
}

func printTrees(trees []*TreeNode) {
	fmt.Print("[")
	for _, tree := range trees {
		fmt.Print("[")
		printTree(tree)
		fmt.Print("], ")
	}
	fmt.Println("]")
}

func printTree(node *TreeNode) {
	if node == nil {
		fmt.Print(node, ", ")
		return
	}
	fmt.Print(node.Val, ", ")
	printTree(node.Left)
	printTree(node.Right)
}
