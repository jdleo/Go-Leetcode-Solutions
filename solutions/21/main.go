package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummy := dummyHead

	// keep going while both list nodes have values
	for l1 != nil && l2 != nil {
		// compare
		if l1.Val < l2.Val {
			// take l1 val and shift
			dummy.Next = &ListNode{Val: l1.Val}
			dummy = dummy.Next
			l1 = l1.Next
		} else {
			// take l2 val and shift
			dummy.Next = &ListNode{Val: l2.Val}
			dummy = dummy.Next
			l2 = l2.Next
		}
	}

	// check if l1 has anything left
	for l1 != nil {
		dummy.Next = &ListNode{Val: l1.Val}
		dummy = dummy.Next
		l1 = l1.Next
	}
	// check if l2 has anything left
	for l2 != nil {
		dummy.Next = &ListNode{Val: l2.Val}
		dummy = dummy.Next
		l2 = l2.Next
	}

	return dummyHead.Next
}
