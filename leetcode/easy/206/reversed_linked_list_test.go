package _206

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//sl := make([]*ListNode, 0)
	sl := make([]int, 0)
	for head != nil {
		sl = append(sl, head.Val)
		head = head.Next
	}

	if len(sl) == 0 {
		return nil
	}
	root := &ListNode{
		Val:  sl[len(sl)-1],
		Next: nil,
	}
	curNode := root
	for i := len(sl) - 2; i >= 0; i-- {
		curNode.Next = &ListNode{
			Val:  sl[i],
			Next: nil,
		}
		curNode = curNode.Next
	}
	return root
}

func TestReversedLinkedList(t *testing.T) {
	tescCases := []struct {
		name string
		head *ListNode
		want string
	}{
		{
			name: "1",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			want: "321",
		},
	}

	for _, tc := range tescCases {
		t.Run(tc.name, func(t *testing.T) {
			root := reverseList(tc.head)
			res := ""
			for root != nil {
				res += strconv.Itoa(root.Val)
				root = root.Next
			}
			assert.Equal(t, tc.want, res)
		})
	}
}
