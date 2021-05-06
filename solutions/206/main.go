package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	// go while there is something in head
	for head != nil {
		// swap
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	// result will be at prev
	return prev
}
