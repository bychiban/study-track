package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list ListNode) Println() {
	for {
		fmt.Printf("%d ", list.Val)
		if list.Next == nil {
			break
		}
		list = *list.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	addTwoNumbers(l1, l2).Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	result := &ListNode{}
	p := result
	for {
		p.Val = p.Val
		if p1 != nil {
			p.Val = p.Val + p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			p.Val = p.Val + p2.Val
			p2 = p2.Next
		}
		if p.Val > 9 {
			p.Val = p.Val - 10
			p.Next = &ListNode{Val: 1}
		} else {
			p.Next = &ListNode{}
		}
		if p1 == nil && p2 == nil {
			if p.Next.Val == 0 {
				p.Next = nil
			}
			break
		}
		p = p.Next
	}
	return result
}
