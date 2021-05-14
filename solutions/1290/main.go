package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0

	// go through linkedlist
	for head != nil {
		// shift over res
		res <<= 1
		// add this bit to res
		res |= head.Val
	}

	return res
}
