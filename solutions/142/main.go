package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	// go while fast can keep moving
	for fast.Next != nil && fast.Next.Next != nil {
		// move slow by 1, fast by 2
		slow = slow.Next
		fast = fast.Next.Next

		// check if we encountered a cycle
		if slow == fast {
			// reset slow
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	// no cycle found
	return nil
}
