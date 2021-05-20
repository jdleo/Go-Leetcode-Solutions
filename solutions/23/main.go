package main

import "math"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// count of lists that have not reached the end
	remainder := len(lists) - nilNodeCount(lists)
	// dummy and head node for result
	head := &ListNode{-1, nil}
	dummy := head

	// go until no more lists to check
	for remainder > 0 {
		// get the minimum node
		min, idx := findMinimum(lists)

		// add this one to result
		dummy.Next = &ListNode{min.Val, nil}
		dummy = dummy.Next

		// move min to next
		lists[idx] = min.Next

		// check if min is now nil
		if lists[idx] == nil {
			remainder--
		}
	}

	return head.Next
}

// find number of nil nodes
func nilNodeCount(lists []*ListNode) int {
	res := 0
	for _, node := range lists {
		if node == nil {
			res++
		}
	}
	return res
}

// helper function to get minimum valued node
func findMinimum(lists []*ListNode) (*ListNode, int) {
	// current minimum
	min := &ListNode{math.MaxInt64, nil}
	minIndex := -1

	// go through nodes
	for i, node := range lists {
		// check if smaller than min
		if node != nil && node.Val < min.Val {
			min = node
			minIndex = i
		}
	}

	return min, minIndex
}
