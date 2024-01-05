package __old

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

func TestRun83(t *testing.T) {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val: 2,
	//				Next: &ListNode{
	//					Val: 2,
	//					Next: &ListNode{
	//						Val: 3,
	//						Next: &ListNode{
	//							Val:  3,
	//							Next: nil,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	//for newHead := deleteDuplicates(head); newHead != nil; newHead = newHead.Next {
	//	fmt.Printf("-> %d\n", newHead.Val)
	//}

	testCases := []struct {
		name string
		head *ListNode
		want string
	}{
		{name: "1", head: nil, want: ""},
		{name: "2", want: "0", head: &ListNode{
			Val:  0,
			Next: nil,
		}},
		{name: "3", want: "123", head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		}},
		{name: "4", want: "123", head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val: 3,
										Next: &ListNode{
											Val: 4,
											Next: &ListNode{
												Val:  4,
												Next: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := deleteDuplicates(tc.head)
			resString := ""
			for head != nil {
				resString += strconv.Itoa(head.Val)
				head = head.Next
			}
			assert.Equal(t, tc.want, resString)
		})
	}

}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}

	return head
}
