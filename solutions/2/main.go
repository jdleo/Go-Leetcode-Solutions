package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummy := dummyHead
	carry := 0

	// go while l1 or l2 has nodes
	for l1 != nil || l2 != nil {
		// get values
		v1, v2 := getVal(l1), getVal(l2)
		// get sum, carry
		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10
		// add to linked list
		dummy.Next = &ListNode{sum, nil}
		dummy = dummy.Next
		// move forward
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// check if something is still in carry
	if carry != 0 {
		// add it to the end of our LL
		dummy.Next = &ListNode{carry, nil}
	}

	return dummyHead.Next
}

// helper method to get node val
func getVal(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}
