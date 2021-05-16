package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	// pre and post (where to insert our list between)
	head := list1
	var pre *ListNode
	var post *ListNode

	// go up to b
	for i := 0; i < b+1; i++ {
		// check if we found the node before a
		if i == a-1 {
			pre = list1
		}

		// check if we found b
		if i == b {
			post = list1.Next
		}

		list1 = list1.Next
	}

	// connect pre to list
	pre.Next = list2
	// find the end of the list
	for list2.Next != nil {
		list2 = list2.Next
	}

	// connect list2's next to post
	list2.Next = post

	return head
}
