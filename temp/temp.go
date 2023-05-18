package main

import "fmt"

func main() {
	fmt.Println(pairSum(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	max := 0
	current := head
	stack Stack := New()
	for current != nil {

	}
}
