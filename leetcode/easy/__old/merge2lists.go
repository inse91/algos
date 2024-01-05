package __old

import "fmt"

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) Add(val int) {

	if ll.Head == nil {
		ll.Head = &ListNode{
			Val:  val,
			Next: nil,
		}
		return
	}
	curNode := ll.Head
	for curNode.Next != nil {
		curNode = curNode.Next
	}

	curNode.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	return
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	ll := NewLinkedList()

	for list1 != nil || list2 != nil {

		if list1 == nil {
			ll.Add(list2.Val)
			list2 = list2.Next
			continue
		}
		if list2 == nil {
			ll.Add(list1.Val)
			list2 = list1.Next
			continue
		}

		if list2.Val <= list1.Val {
			ll.Add(list2.Val)
			list2 = list2.Next
		} else {
			ll.Add(list1.Val)
			list1 = list1.Next
		}

	}

	return ll.Head
}

func mergeTwoListsNew(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsNew(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoListsNew(list1, list2.Next)
	return list2

}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) (firstOut *ListNode) {
//
//	out := make([]int, 0)
//
//	for list1 != nil || list2 != nil {
//
//		if list1 == nil {
//			out = append(out, list2.Val)
//			list2 = list2.Next
//			continue
//		}
//		if list2 == nil {
//			out = append(out, list1.Val)
//			list2 = list1.Next
//			continue
//		}
//
//		var min int
//		if list2.Val <= list1.Val {
//			min = list2.Val
//			list2 = list2.Next
//		} else {
//			min = list1.Val
//			list1 = list1.Next
//		}
//		out = append(out, min)
//
//	}
//	fmt.Printf("%+v", out)
//	if len(out) == 0 {
//		return
//	}
//
//	firstOut = &ListNode{
//		Val:  out[0],
//		Next: nil,
//	}
//	out = out[1:]
//
//	curNode := firstOut
//	for _, num := range out {
//		curNode.Next = &ListNode{
//			Val:  num,
//			Next: nil,
//		}
//		curNode = curNode.Next
//	}
//	return
//}

func RunMerge2Lists() {
	node := mergeTwoLists(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  6,
					Next: nil,
				},
			},
		},
	)

	for node != nil {
		fmt.Printf(" -> %d", node.Val)
		node = node.Next
	}
}
