package main

import "fmt"

func main() {
	fmt.Println(swapNodes(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	count := 0
	current := head
	pFirst := head
	pSecond := head
	startMove := false
	for current != nil {
		count++
		if startMove {
			pSecond = pSecond.Next
		}
		if count == k {
			pFirst = current
			startMove = true
		}
		current = current.Next
	}

	pFirst.Val, pSecond.Val = pSecond.Val, pFirst.Val
	return head
}
