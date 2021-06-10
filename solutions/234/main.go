package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	// first, find the middle node
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// next, reverse the second half
	var prev *ListNode = nil
	for slow != nil {
		// 1 -> 2 -> 3
		// null <- 1
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// lastly, compare head with second half
	for prev != nil {
		// compare with prev (which is head of second half)
		if head.Val != prev.Val {
			return false
		}
		prev, head = prev.Next, head.Next
	}

	// no issues
	return true
}
