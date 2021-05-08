package main

import "math"

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// stack for inorder
	stack := []*TreeNode{}
	// previous val (min int to start)
	prev := math.MinInt64

	// go while stack or root is not nil
	for len(stack) > 0 || root != nil {
		if root != nil {
			// add to stack, and keep going left
			stack = append(stack, root)
			root = root.Left
		} else {
			// pop from stack
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// make sure this is strictly greater than last
			if root.Val <= prev {
				return false
			}
			// set new prev, and now go right
			prev = root.Val
			root = root.Right
		}
	}

	// no issues
	return true
}
