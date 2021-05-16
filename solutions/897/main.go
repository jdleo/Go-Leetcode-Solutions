package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	dummy := head

	// stack for inorder
	stack := []*TreeNode{}
	// go while stack or root
	for len(stack) > 0 || root != nil {
		// check root
		if root != nil {
			// keep goin left
			stack = append(stack, root)
			root = root.Left
		} else {
			// pop from stack
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// add to result (INORDER)
			dummy.Right = root
			root.Left = nil
			dummy = dummy.Right
			// go right
			root = root.Right
		}
	}

	return head.Right
}
