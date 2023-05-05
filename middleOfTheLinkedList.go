package main

import (
	"fmt"
)

func main() {
	fmt.Println(middleNode(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{Val: 5}}}}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	result := head
	for current := head; current != nil && current.Next != nil; current = current.Next.Next {
		result = result.Next
	}
	return result
}
