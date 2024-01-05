package __old

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/description/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	acc, nd := 0, new(ListNode)
	for node := nd; l1 != nil || l2 != nil || acc > 0; node = node.Next {
		if l1 != nil {
			acc += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			acc += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{
			Val:  acc % 10,
			Next: nil,
		}
		acc = acc / 10

	}
	return nd.Next
}

func RunTwoNumbers() {
	node := addTwoNumbers(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	)

	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}
