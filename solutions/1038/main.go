package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	// stack for inorder traversal, and running sum
	stack, sum, cur := []*TreeNode{}, 0, root

	// go while nodes to traverse
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			// add to stack (but go right, REVERSED inorder)
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			// pop from stack
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// add to running sum, and assign to this node (will have all nodes vals before it)
			sum += cur.Val
			cur.Val = sum
			// now go left
			cur = cur.Left
		}
	}

	return root
}
