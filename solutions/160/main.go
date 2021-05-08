package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// pointers to traverse
	a, b := headA, headB

	// go until they meet
	for a != b {
		// move both forward, but if we reach end, reset back to head
		if a != nil {
			a = a.Next
		} else {
			a = headA
		}
		if b != nil {
			b = b.Next
		} else {
			b = headB
		}
	}

	return a
}
